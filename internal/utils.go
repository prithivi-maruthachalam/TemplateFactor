package internal

import (
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

type Template struct {
	TemplateName string
	Nodes        []TemplateNode
}
