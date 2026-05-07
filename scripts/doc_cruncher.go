package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	providerName := flag.String("provider", "", "Name of the provider")
	docsDir := flag.String("docs", "", "Path to the provider's resource documentation directory")
	flag.Parse()

	if *providerName == "" || *docsDir == "" {
		fmt.Println("Usage: go run . -provider=<name> -docs=<path>")
		os.Exit(1)
	}

	files, err := filepath.Glob(filepath.Join(*docsDir, "*.md"))
	if len(files) == 0 {
		files, err = filepath.Glob(filepath.Join(*docsDir, "*.html.markdown"))
	}
	if err != nil || len(files) == 0 {
		fmt.Printf("No markdown files found in %s\n", *docsDir)
		os.Exit(1)
	}

	strategies := map[string]ProviderStrategy{
		"aws":        getAWSStrategy(),
		"kubernetes": getKubernetesStrategy(),
		"argocd":     getArgocdStrategy(),
		"google":     getGoogleStrategy(),
		"azurerm":    getAzurermStrategy(),
		"scaleway":   getScalewayStrategy(),
	}

	strategy, exists := strategies[strings.ToLower(*providerName)]
	if !exists {
		fmt.Printf("Warning: Provider '%s' does not have a dedicated strategy. Falling back to AWS strategy.\n", *providerName)
		strategy = strategies["aws"]
	}

	var mappings []mapping

	for _, file := range files {
		base := filepath.Base(file)
		resourceName := fmt.Sprintf("%s_%s", *providerName, strings.TrimSuffix(strings.TrimSuffix(base, ".html.markdown"), ".md"))
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		match := strategy.MatchRegex.FindStringSubmatch(string(content))
		m := mapping{Resource: resourceName}

		if len(match) > 1 {
			keys, isComputed, comment := strategy.ExtractFunc(match)
			m.Keys = keys
			m.IsComputed = isComputed
			m.Comment = comment
		} else {
			m.IsComputed = true
			m.Comment = "No standard import format found in documentation"
		}
		mappings = append(mappings, m)
	}

	sort.Slice(mappings, func(i, j int) bool {
		return mappings[i].Resource < mappings[j].Resource
	})

	funcName := fmt.Sprintf("extract%sImportID", strings.ToUpper((*providerName)[:1])+(*providerName)[1:])
	if strings.ToLower(*providerName) == "aws" {
		funcName = "extractAWSImportID"
	}

	generatedCode := strategy.GenerateFunc(mappings, funcName, *providerName)

	formatted, err := format.Source([]byte(generatedCode))
	if err != nil {
		fmt.Printf("Error formatting generated code: %v\n", err)
		os.Exit(1)
	}

	outPath := filepath.Join("..", "internal", "providers", fmt.Sprintf("%s.go", *providerName))
	err = os.WriteFile(outPath, formatted, 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outPath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s\n", outPath)
}
