package simple_factory

import (
	"fmt"
)

// ConfigFileParser
type ConfigFileParser interface {
	Parse(filename string)
}

// YamlConfigParser
type YamlConfigParser struct{}

func (y YamlConfigParser) Parse(filename string) {
	fmt.Printf("%s yaml file parser", filename)
}

// JsonConfigParser
type JsonConfigParser struct{}

func (y JsonConfigParser) Parse(filename string) {
	fmt.Printf("%s yaml file parser", filename)
}

// NewConfigFileParser
func NewConfigFileParser(parser string, filename string) ConfigFileParser {
	switch parser {
	case "json":
		return JsonConfigParser{}
	case "yaml":
		return YamlConfigParser{}
	}
	return nil
}
