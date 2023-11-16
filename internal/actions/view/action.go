package view

import (
	"fmt"
	"log"

	tf_errors "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
	tf_storage "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/storage"
)

func ViewTemplate(templateName string, showContent bool) {
	template, err := tf_storage.LoadTemplate(templateName)
	if template == nil && err == nil {
		log.Fatal(&tf_errors.TemplateNotFoundError{TemplateName: templateName})
	} else if err != nil {
		log.Fatal(&tf_errors.InternalError{Cause: err, Name: tf_errors.StorageError})
	}

	fmt.Print(template.Describe(showContent))
}
