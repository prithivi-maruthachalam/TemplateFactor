package actions

import (
	"fmt"
)

type CreateTemplateConfig struct {
	TemplateName    string   // Name of the template
	SourceDirPath   string   // The source directory for the template
	SaveFiles       bool     `default:"false"` // Should files be included in the template
	SaveFileContent bool     `default:"false"` // Should the content of files be included in the template
	StoreLinks      bool     `default:"false"` // Should dirs and contents be stored as links
	Clobber         bool     `default:"false"` // Should an existing template with the same name be overwritten
	IgnoreList      []string // List of glob patterns to ignore from the template
}

func CreateTemplate(params CreateTemplateConfig) {
	fmt.Println("Create Template")
	fmt.Print(params)
}
