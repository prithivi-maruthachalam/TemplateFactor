package delete

import (
	"log"

	tf_errors "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
	tf_storage "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/storage"
)

func DeleteTemplate(templateName string) {
	err := tf_storage.DeleteTemplate(templateName)
	if err != nil {
		log.Fatal(&tf_errors.InternalError{Cause: err, Name: tf_errors.StorageError})
	}
}
