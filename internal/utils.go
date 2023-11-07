package internal

import (
	"fmt"
	"log"
	"regexp"

	"github.com/bmatcuk/doublestar/v4"
)

// Validates the name of the template
func ValidateTemplateName(name string) bool {
	match, err := regexp.MatchString("^[a-zA-Z$][a-zA-Z_$0-9]+$", name)

	if err != nil {
		log.Fatal(err)
	}

	return match
}

// Test if the given path string matches any of the glob patterns in the slice
func TestGlobMatches(patterns []string, path string) (bool, error) {
	for _, pattern := range patterns {
		match, err := doublestar.Match(pattern, path)
		if err != nil {
			return false, err
		}

		if match {
			return match, nil
		}
	}

	return false, nil
}

// Represents a single node in a template. This can be a directory or a file with or without content
type TemplateNode struct {
	NodePath string
	IsFile   bool
	Content  string
}

func (node *TemplateNode) String() string {
	nodeTypeStr := "dir"
	if node.IsFile {
		nodeTypeStr = "file"
	}
	return fmt.Sprintf("  - %s %s content:%d\n", node.NodePath, nodeTypeStr, len(node.Content))
}

// Template factory template object
type Template struct {
	TemplateName string
	Nodes        []TemplateNode
}

// Add a file/directory node to the template
func (template *Template) AddNode(node TemplateNode) {
	template.Nodes = append(template.Nodes, node)
}

func (template *Template) String() string {
	str := fmt.Sprintf("Name: %s\n", template.TemplateName)

	for _, node := range template.Nodes {
		str += fmt.Sprintf("%v", &node)
	}

	return str
}
