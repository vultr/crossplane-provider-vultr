package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/alecthomas/kingpin/v2"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/pipeline"

	"github.com/crossplane-contrib/provider-vultr/config"
)

func main() {
	var (
		app                   = kingpin.New("generator", "Run Upjet code generation pipelines for provider-vultr").DefaultEnvars()
		repoRoot              = app.Arg("repo-root", "Root directory for the provider repository").Required().String()
		generatedResourceList = app.Flag("generated-resource-list", "File path where a list of the generated resources will be stored.").Envar("GENERATED_RESOURCE_LIST").Default("../config/generated.lst").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	absRootDir, err := filepath.Abs(*repoRoot)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", *repoRoot))
	}
	pc := config.GetProvider()
	dumpGeneratedResourceList(pc, *generatedResourceList)
	pipeline.Run(pc, config.GetProviderNamespaced(), absRootDir)
}

// dumpGeneratedResourceList writes the sorted Terraform names of the
// generated resources, used by `make schema-version-diff`.
func dumpGeneratedResourceList(p *ujconfig.Provider, targetPath string) {
	if targetPath == "" {
		return
	}
	generatedResources := make([]string, 0, len(p.Resources))
	for name := range p.Resources {
		generatedResources = append(generatedResources, name)
	}
	sort.Strings(generatedResources)
	buff, err := json.MarshalIndent(generatedResources, "", "")
	if err != nil {
		panic(fmt.Sprintf("cannot marshal the generated resource list to JSON: %s", err.Error()))
	}
	if err := os.WriteFile(targetPath, buff, 0o600); err != nil {
		panic(fmt.Sprintf("cannot write the generated resource list to file %s: %s", targetPath, err.Error()))
	}
}
