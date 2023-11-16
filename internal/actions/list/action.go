package list

import (
	"fmt"
	"log"

	tf_errors "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
	tf_storage "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/storage"
)

func ListTemplate() {
	templates, err := tf_storage.GetAllTemplates()
	if err != nil {
		log.Fatal(&tf_errors.InternalError{Cause: err, Name: tf_errors.StorageError})
	}

	for _, template := range templates {
		fmt.Println(tf_io.TemplateName(template))
	}
}
