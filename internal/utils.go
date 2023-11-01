package internal

import (
	"fmt"
	"log"
	"regexp"
)

func ValidateTemplateName(name string) bool {
	match, err := regexp.MatchString("^[a-zA-Z$][a-zA-Z_$0-9]+$", name)

	if err != nil {
		log.Fatal(err)
	}

	return match
}

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
	return fmt.Sprintf("    - %s %s %s\n", node.NodePath, nodeTypeStr, node.Content)
}

type Template struct {
	TemplateName string
	Nodes        []TemplateNode
}

func (template *Template) AddNode(node TemplateNode) {
	template.Nodes = append(template.Nodes, node)
}
