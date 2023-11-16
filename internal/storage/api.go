package storage

import (
	"os"
	"path/filepath"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common"
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/utils"
)

var homedir = utils.GetUserHomeDir()
var TF_HOME = filepath.Join(homedir, ".templatefactory")

const home_permissions = 0777

func TestAndCreateTemplateFactoryHome() error {
	return os.MkdirAll(TF_HOME, home_permissions)
}

func StoreTemplate(template *common.Template) error {
	encodedTemplate, err := encodeToGOB64(template)
	if err != nil {
		return err
	}

	err = save(template.TemplateName, encodedTemplate)
	if err != nil {
		return err
	}

	return nil
}

func LoadTemplate(templateName string) (*common.Template, error) {
	encodedTemplate, err := load(templateName)
	if err != nil {
		return nil, err
	}

	if encodedTemplate == "" {
		return nil, nil
	}

	decodedTemplate, err := decodeFromGOB64(encodedTemplate)
	if err != nil {
		return nil, nil
	}

	return &decodedTemplate, nil

}
