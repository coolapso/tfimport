package main

import "regexp"

type mapping struct {
	Resource      string
	Keys          []string
	IsComputed    bool
	Comment       string
	FormatStrings []string
}

type ProviderStrategy struct {
	MatchRegex   *regexp.Regexp
	ExtractFunc  func(match []string) ([]string, bool, string)
	GenerateFunc func(mappings []mapping, funcName, providerName string) string
}
