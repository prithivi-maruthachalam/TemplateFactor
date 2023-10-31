package actions

import (
	"fmt"
)

type CreateTemplateConfig struct {
	TemplateName       string
	SourceDirPath      string `default:"."`
	SaveFiles          bool   `default:"false"`
	SaveContent        bool   `default:"false"`
	StoreLink          bool   `default:"false"`
	Clobber            bool   `default:"false"`
	DryRun             bool   `default:"false"`
	ExcludeList        []string
	FileIncludeList    []string
	ContentExcludeList []string
	ContentIncludeList []string
}

func CreateTemplate(params CreateTemplateConfig) {
	fmt.Println("Create Template")
	fmt.Print(params)
}
