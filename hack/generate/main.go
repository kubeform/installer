/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	goflag "flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	flag "github.com/spf13/pflag"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crdv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/parser"
	"sigs.k8s.io/yaml"
)

var (
	crdstore   = map[schema.GroupKind]map[string]*unstructured.Unstructured{}
	extrastore = map[schema.GroupKind]map[string]*unstructured.Unstructured{}
	catalog    = map[string][]string{}
)

type Definition struct {
	Kind       string         `json:"kind,omitempty"`
	APIVersion string         `json:"apiVersion,omitempty"`
	Spec       DefinitionSpec `json:"spec"`
}

type DefinitionSpec struct {
	Group string                              `json:"group"`
	Names crdv1.CustomResourceDefinitionNames `json:"names"`
}

type ProviderList struct {
	Providers []string
}

type CRD struct {
	GK  schema.GroupKind
	Def []byte
}

var providerList ProviderList

type ProviderData struct {
	Provider    string
	Groups      []string
	GIDs        []string
	OperatorTag string
}

type ProviderGroupData struct {
	ProviderData
	GID string
}

/*
.Providers ===> lower
.Provider  ===> lower
.Groups    ===> full group
.GIDs      ===> group prefix (lower)
.GID       ===> lower
*/

func main() {
	var providers = []string{
		"aws",
		"azurerm",
		"digitalocean",
		"google",
		"linode",
	}
	var provider string
	var gid string
	var inputDir string = "/home/tamal/go/src/kubeform.dev"
	var extraInput = []string{
		"https://github.com/kmodules/custom-resources/raw/kubernetes-1.21.1/crds/metrics.appscode.com_metricsconfigurations.v1.yaml",
	}
	var crdVersion = "v1"

	flag.StringSliceVar(&providers, "providers", providers, "List of providers")
	flag.StringVar(&provider, "provider", provider, "Provider to be processed, if empty process all providers")
	flag.StringVar(&gid, "group-id", gid, "Only process group id for --provider if set")
	flag.StringVar(&inputDir, "input-dir", inputDir, "Directory which contains provider-***-api directories")
	flag.StringSliceVar(&extraInput, "extra-input", extraInput, "List of extra crd urls or dir/files")
	flag.StringVar(&crdVersion, "v", crdVersion, "CRD version v1/v1beta1")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	for _, location := range extraInput {
		err := processLocation(location, extrastore)
		if err != nil {
			panic(err)
		}
	}

	strset := sets.NewString()
	for _, p := range providers {
		strset.Insert(strings.ToLower(strings.TrimSpace(p)))
	}
	providerList = ProviderList{
		Providers: strset.List(),
	}
	provider = strings.ToLower(strings.TrimSpace(provider))

	catalogfile := filepath.Join(inputDir, "installer", "crds", "catalog.json")
	err := os.MkdirAll(filepath.Dir(catalogfile), 0755)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(catalogfile); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(catalogfile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &catalog)
		if err != nil {
			panic(err)
		}
	}

	{
		providersToProcess := strset.List()
		if provider != "" {
			providersToProcess = []string{provider}
		}
		for _, p := range providersToProcess {
			err := processProvider(inputDir, p, gid, crdVersion)
			if err != nil {
				panic(err)
			}
		}
	}

	data, err := json.MarshalIndent(catalog, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(catalogfile, data, 0644)
	if err != nil {
		panic(err)
	}
}

func processProvider(inputDir string, p, gid, crdVersion string) error {
	crdstore = map[schema.GroupKind]map[string]*unstructured.Unstructured{} // reset crd store

	location := filepath.Join(inputDir, fmt.Sprintf("provider-%s-api", p), "crds")
	err := processLocation(location, crdstore)
	if err != nil {
		return err
	}

	groups := sets.NewString()
	gids := sets.NewString()
	for gk := range crdstore {
		groups.Insert(gk.Group)
		gids.Insert(gk.Group[0:strings.IndexRune(gk.Group, '.')])
	}
	catalog[p] = groups.List()

	err = os.MkdirAll(filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s", p)), 0755)
	if err != nil {
		return err
	}

	// remove crd charts for api groups that don't exist anymore
	entries, err := ioutil.ReadDir(filepath.Join(inputDir, "installer", "charts"))
	if err != nil {
		return err
	}
	for _, fi := range entries {
		if fi.IsDir() &&
			strings.HasPrefix(fi.Name(), fmt.Sprintf("kubeform-provider-%s-", p)) &&
			!gids.Has(strings.TrimSuffix(strings.TrimPrefix(fi.Name(), fmt.Sprintf("kubeform-provider-%s-", p)), "-crds")) {
			_ = os.RemoveAll(filepath.Join(inputDir, "installer", "charts", fi.Name()))
		}
	}

	var buf bytes.Buffer

	// installer/.generator/.github/workflows/
	fmt.Println(filepath.Join(inputDir, "installer", ".github", "workflows"))
	root := filepath.Join(inputDir, "installer", ".generator", ".github", "workflows")
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		/*
			.
			ci.yml
		*/

		target := filepath.Join(inputDir, "installer", ".github", "workflows", rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(path))
		buf.Reset()
		err = tpl.ExecuteTemplate(&buf, filepath.Base(path), providerList)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(target, buf.Bytes(), 0644)
	})
	if err != nil {
		return err
	}

	// installer/.generator/apis
	fmt.Println(filepath.Join(inputDir, "installer", "apis"))
	root = filepath.Join(inputDir, "installer", ".generator", "apis")
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		/*
			.
			installer
			installer/fuzzer
			installer/fuzzer/fuzzer.go
			installer/v1alpha1
			installer/v1alpha1/kubeform_provider_p.go
			installer/v1alpha1/register.go
			installer/v1alpha1/types_test.go
		*/

		target := filepath.Join(inputDir, "installer", "apis", rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		switch rel {
		case "installer/v1alpha1/kubeform_provider_p.go":
			tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(path))
			buf.Reset()
			err = tpl.ExecuteTemplate(&buf, filepath.Base(path), ProviderData{
				Provider: p,
				Groups:   groups.List(),
				GIDs:     gids.List(),
			})
			if err != nil {
				return err
			}
			return ioutil.WriteFile(filepath.Join(filepath.Dir(target), fmt.Sprintf("kubeform_provider_%s.go", p)), buf.Bytes(), 0644)
		default:
			tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(path))
			buf.Reset()
			err = tpl.ExecuteTemplate(&buf, filepath.Base(path), providerList)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(target, buf.Bytes(), 0644)
		}
	})
	if err != nil {
		return err
	}

	// installer/.generator/charts/kubeform-provider-p
	fmt.Println(filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s", p)))
	root = filepath.Join(inputDir, "installer", ".generator", "charts", "kubeform-provider-p")
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		/*
			.
			.helmignore
			Chart.yaml
			ci
			ci/ci-values.yaml
			doc.yaml
			templates
			templates/NOTES.txt
			templates/user-roles.yaml
			values.yaml
		*/

		target := filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s", p), rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		switch rel {
		case "templates/NOTES.txt":
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(target, data, 0644)
		default:
			operatorTag := "v0.1.0" //default tag
			if rel == "values.yaml" {
				if _, e2 := os.Stat(target); !os.IsNotExist(e2) {
					// keep original kubeform-provider.operator.tag

					var existing map[string]interface{}
					data, err := ioutil.ReadFile(target)
					if err != nil {
						return err
					}
					err = yaml.Unmarshal(data, &existing)
					if err != nil {
						return err
					}

					tag, exists, err := unstructured.NestedString(existing, "kubeform-provider", "operator", "tag")
					if err != nil {
						return err
					}
					if exists {
						operatorTag = tag
					}
				}
			}

			tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(path))
			buf.Reset()
			err = tpl.ExecuteTemplate(&buf, filepath.Base(path), ProviderData{
				Provider:    p,
				Groups:      groups.List(),
				GIDs:        gids.List(),
				OperatorTag: operatorTag,
			})
			if err != nil {
				return err
			}
			if rel == "Chart.yaml" {
				if _, e2 := os.Stat(target); !os.IsNotExist(e2) {
					// keep original version, appversion

					var existing map[string]interface{}
					data, err := ioutil.ReadFile(target)
					if err != nil {
						return err
					}
					err = yaml.Unmarshal(data, &existing)
					if err != nil {
						return err
					}

					var nu map[string]interface{}
					err = yaml.Unmarshal(buf.Bytes(), &nu)
					if err != nil {
						return err
					}
					if v, ok := existing["version"]; ok && v.(string) != "" {
						nu["version"] = v
					}
					if v, ok := existing["appVersion"]; ok && v.(string) != "" {
						nu["appVersion"] = v
					}
					data, err = yaml.Marshal(nu)
					if err != nil {
						return err
					}
					buf.Reset()
					buf.Write(data)
				}
			}
			return ioutil.WriteFile(target, buf.Bytes(), 0644)
		}
	})
	if err != nil {
		return err
	}

	// installer/.generator/charts/kubeform-provider-p-g-crds

	gidsToProcess := gids.List()
	if gid != "" {
		gidsToProcess = []string{gid}
	}
	for _, g := range gidsToProcess {
		fmt.Println(filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s-%s-crds", p, g)))

		// remove crds folder so that we can cleanly recreate it.
		_ = os.RemoveAll(filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s-%s-crds", p, g), "crds"))

		outputCRDs := make([]CRD, 0, len(crdstore)+len(extrastore))

		root = filepath.Join(inputDir, "installer", ".generator", "charts", "kubeform-provider-p-g-crds")
		err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			/*
				.
				.helmignore
				Chart.yaml
				crds
				crds/metrics.appscode.com_metricsconfigurations.yaml
				doc.yaml
				templates
				templates/NOTES.txt
				templates/metrics-user-roles.yaml
				values.yaml
			*/

			target := filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s-%s-crds", p, g), rel)

			if info.IsDir() {
				return os.MkdirAll(target, 0755)
			}

			switch rel {
			case "templates/NOTES.txt":
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				return ioutil.WriteFile(target, data, 0644)
			default:
				tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(path))
				buf.Reset()
				err = tpl.ExecuteTemplate(&buf, filepath.Base(path), ProviderGroupData{
					ProviderData: ProviderData{
						Provider: p,
						Groups:   groups.List(),
						GIDs:     gids.List(),
					},
					GID: g,
				})
				if err != nil {
					return err
				}
				if rel == "Chart.yaml" {
					if _, e2 := os.Stat(target); !os.IsNotExist(e2) {
						// keep original version, appversion

						var existing map[string]interface{}
						data, err := ioutil.ReadFile(target)
						if err != nil {
							return err
						}
						err = yaml.Unmarshal(data, &existing)
						if err != nil {
							return err
						}

						var nu map[string]interface{}
						err = yaml.Unmarshal(buf.Bytes(), &nu)
						if err != nil {
							return err
						}
						if v, ok := existing["version"]; ok && v.(string) != "" {
							nu["version"] = v
						}
						if v, ok := existing["appVersion"]; ok && v.(string) != "" {
							nu["appVersion"] = v
						}
						data, err = yaml.Marshal(nu)
						if err != nil {
							return err
						}
						buf.Reset()
						buf.Write(data)
					}
				}
				return ioutil.WriteFile(target, buf.Bytes(), 0644)
			}
		})
		if err != nil {
			return err
		}

		// crds
		target := filepath.Join(inputDir, "installer", "charts", fmt.Sprintf("kubeform-provider-%s-%s-crds", p, g), "templates", "crds")
		err = os.MkdirAll(target, 0755)
		if err != nil {
			return err
		}
		for gk := range crdstore {
			if !strings.HasPrefix(gk.Group, g+".") {
				continue
			}
			data, filename, err := WriteCRD(crdstore, target, gk, crdVersion)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(filename, data, 0644)
			if err != nil {
				panic(err)
			}

			outputCRDs = append(outputCRDs, CRD{
				GK:  gk,
				Def: data,
			})
		}
		for gk := range extrastore {
			data, filename, err := WriteCRD(extrastore, target, gk, crdVersion)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(filename, data, 0644)
			if err != nil {
				panic(err)
			}

			outputCRDs = append(outputCRDs, CRD{
				GK:  gk,
				Def: data,
			})
		}

		// write kubeform-provider-p-g-crds.yaml file
		sort.Slice(outputCRDs, func(i, j int) bool {
			if outputCRDs[i].GK.Group == outputCRDs[j].GK.Group {
				return outputCRDs[i].GK.Kind < outputCRDs[j].GK.Kind
			}
			return outputCRDs[i].GK.Group < outputCRDs[j].GK.Group
		})

		target = filepath.Join(inputDir, "installer", "crds", fmt.Sprintf("kubeform-provider-%s", p), fmt.Sprintf("%s-crds.yaml", g))
		err = os.MkdirAll(filepath.Dir(target), 0755)
		if err != nil {
			return err
		}

		buf.Reset()
		for i, crd := range outputCRDs {
			if i > 0 {
				buf.WriteString("\n---\n")
			}
			buf.Write(crd.Def)
		}
		err = ioutil.WriteFile(target, buf.Bytes(), 0644)
		if err != nil {
			return err
		}
	}

	// installer/.generator/hack/scripts
	fmt.Println(filepath.Join(inputDir, "installer", "hack", "scripts"))
	root = filepath.Join(inputDir, "installer", ".generator", "hack", "scripts")
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		/*
			.
			import-crds.sh
			update-chart-dependencies.sh
		*/

		target := filepath.Join(inputDir, "installer", "hack", "scripts", rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(path))
		buf.Reset()
		err = tpl.ExecuteTemplate(&buf, filepath.Base(path), providerList)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(target, buf.Bytes(), 0755)
	})
	if err != nil {
		return err
	}

	return nil
}

func processLocation(location string, store map[schema.GroupKind]map[string]*unstructured.Unstructured) error {
	u, err := url.Parse(location)
	if err != nil {
		return err
	}

	if u.Scheme != "" {
		resp, err := http.Get(u.String())
		if err != nil {
			return err
		}
		defer func() { _ = resp.Body.Close() }()
		var buf bytes.Buffer
		_, err = io.Copy(&buf, resp.Body)
		if err != nil {
			return err
		}
		return parser.ProcessResources(buf.Bytes(), extractCRD(store))
	}

	fi, err := os.Stat(location)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return parser.ProcessDir(location, extractCRD(store))
	} else {
		data, err := ioutil.ReadFile(location)
		if err != nil {
			return err
		}
		return parser.ProcessResources(data, extractCRD(store))
	}
}

func extractCRD(store map[schema.GroupKind]map[string]*unstructured.Unstructured) func(obj *unstructured.Unstructured) error {
	return func(obj *unstructured.Unstructured) error {
		removeDescription(obj.UnstructuredContent())

		obj.SetAnnotations(map[string]string{
			"helm.sh/resource-policy": "keep",
		})

		var def Definition

		err := meta_util.DecodeObject(obj.Object, &def)
		if err != nil {
			return err
		}

		gv, err := schema.ParseGroupVersion(def.APIVersion)
		if err != nil {
			return err
		}

		gk := schema.GroupKind{
			Group: def.Spec.Group,
			Kind:  def.Spec.Names.Kind,
		}

		if _, ok := store[gk]; !ok {
			store[gk] = map[string]*unstructured.Unstructured{}
		}
		store[gk][gv.Version] = obj

		return nil
	}
}

func WriteCRD(store map[schema.GroupKind]map[string]*unstructured.Unstructured, dir string, gk schema.GroupKind, version string) ([]byte, string, error) {
	crdversions, ok := store[gk]
	if !ok {
		return nil, "", fmt.Errorf("missing crd for %+v", gk)
	}
	if len(crdversions) == 0 {
		return nil, "", fmt.Errorf("missing crd version for %+v", gk)
	}

	crd, ok := crdversions[version]
	if !ok {
		if version == "v1" {
			// convert to v1
			data, err := yaml.Marshal(crdversions["v1beta1"])
			if err != nil {
				return nil, "", err
			}
			var defv1beta1 crdv1beta1.CustomResourceDefinition
			err = yaml.Unmarshal(data, &defv1beta1)
			if err != nil {
				return nil, "", err
			}

			var inner apiextensions.CustomResourceDefinition
			err = crdv1beta1.Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(&defv1beta1, &inner, nil)
			if err != nil {
				return nil, "", err
			}

			var defv1 crdv1.CustomResourceDefinition
			err = crdv1.Convert_apiextensions_CustomResourceDefinition_To_v1_CustomResourceDefinition(&inner, &defv1, nil)
			if err != nil {
				return nil, "", err
			}

			data, err = yaml.Marshal(defv1)
			if err != nil {
				return nil, "", err
			}

			filename := filepath.Join(dir, fmt.Sprintf("%s_%s.yaml", defv1.Spec.Group, defv1.Spec.Names.Plural))
			return data, filename, nil
			// return ioutil.WriteFile(filename, data, 0644)
		} else if version == "v1beta1" {
			// convert to v1beta1
			data, err := yaml.Marshal(crdversions["v1"])
			if err != nil {
				return nil, "", err
			}
			var defv1 crdv1.CustomResourceDefinition
			err = yaml.Unmarshal(data, &defv1)
			if err != nil {
				return nil, "", err
			}

			var inner apiextensions.CustomResourceDefinition
			err = crdv1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(&defv1, &inner, nil)
			if err != nil {
				return nil, "", err
			}

			var defv1beta1 crdv1beta1.CustomResourceDefinition
			err = crdv1beta1.Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(&inner, &defv1beta1, nil)
			if err != nil {
				return nil, "", err
			}

			data, err = yaml.Marshal(defv1beta1)
			if err != nil {
				return nil, "", err
			}

			filename := filepath.Join(dir, fmt.Sprintf("%s_%s.yaml", defv1beta1.Spec.Group, defv1beta1.Spec.Names.Plural))
			return data, filename, nil
		}
	}

	data, err := yaml.Marshal(crd)
	if err != nil {
		return nil, "", err
	}

	var def Definition
	err = meta_util.DecodeObject(crd.Object, &def)
	if err != nil {
		return nil, "", err
	}
	filename := filepath.Join(dir, fmt.Sprintf("%s_%s.yaml", def.Spec.Group, def.Spec.Names.Plural))
	return data, filename, nil
}

func removeDescription(m map[string]interface{}) {
	for k, v := range m {
		if k == "description" {
			if _, ok := v.(string); ok {
				delete(m, k)
			}
		}
		if inner, ok := v.(map[string]interface{}); ok {
			removeDescription(inner)
		} else if arr, ok := v.([]interface{}); ok {
			for i := range arr {
				if inner, ok := arr[i].(map[string]interface{}); ok {
					removeDescription(inner)
				}
			}
		}
	}
}
