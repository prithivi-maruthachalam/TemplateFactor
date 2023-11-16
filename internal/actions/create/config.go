package create

import (
	"fmt"
	"os"
	"strings"

	tf_errors "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
	tf_utils "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/utils"
)

// Contains the input configuration used to create a template
type CreateTemplateConfig struct {
	TemplateName       string
	SourceDirPath      string
	SaveFiles          bool
	SaveContent        bool
	Clobber            bool
	DryRun             bool
	ExcludeList        []string
	FileIncludeList    []string
	ContentExcludeList []string
	ContentIncludeList []string
}

// Stringer implementation for CreateTemplateConfig
func (config *CreateTemplateConfig) String() string {
	str := fmt.Sprintf("\n  Template Name: %s\n  Source Dir: %s\n  SaveFile?: %t\n  SaveContent?: %t\n  Clobber?: %t\n  DryRun?: %t\n",
		config.TemplateName,
		config.SourceDirPath,
		config.SaveFiles,
		config.SaveContent,
		config.Clobber,
		config.DryRun)

	str += "  ExcludeList: ("
	for _, pat := range config.ExcludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += "  FileIncludeList: ("
	for _, pat := range config.FileIncludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += "  ContentExcludeList: ("
	for _, pat := range config.ContentExcludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += "  ContentIncludeList: ("
	for _, pat := range config.ContentIncludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += strings.Repeat("-", 10)
	str += "\n"

	return str
}

// Validates the input configuration. Returns an error if invalid.
func (config *CreateTemplateConfig) Validate() error {
	if !tf_utils.ValidateTemplateName(config.TemplateName) {
		return &tf_errors.InvalidTemplateNameError{TemplateName: config.TemplateName}
	}

	if _, err := os.Stat(config.SourceDirPath); os.IsNotExist(err) {
		return &tf_errors.SourceDirNotFoundError{SourceDir: config.SourceDirPath}
	} else if err != nil {
		return &tf_errors.InternalError{Cause: err, Name: tf_errors.SourceDirStatError}
	}

	return nil
}
