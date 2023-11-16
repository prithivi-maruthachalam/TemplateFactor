package storage

import (
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common"
)

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

func DeleteTemplate(templateName string) error {
	err := delete(templateName)
	return err
}

func GetAllTemplates() ([]string, error) {
	templates, err := getAllKeys()
	if err != nil {
		return nil, nil
	}

	return templates, nil
}

func LoadTemplate(templateName string) (*common.Template, error) {
	encodedTemplate, err := load(templateName)
	if err != nil {
		return nil, err
	}

	if len(encodedTemplate) == 0 {
		return nil, nil
	}

	decodedTemplate, err := decodeFromGOB64(encodedTemplate)
	if err != nil {
		return nil, nil
	}

	return &decodedTemplate, nil

}
