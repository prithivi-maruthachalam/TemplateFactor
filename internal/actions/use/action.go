package use

import (
	"io/fs"
	"log"
	"os"

	tf_errors "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
	tf_storage "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/storage"
)

func UseTemplate(params UseTemplateConfig) {
	// Get the template
	template, err := tf_storage.LoadTemplate(params.TemplateName)
	if template == nil && err == nil {
		log.Fatal(&tf_errors.TemplateNotFoundError{TemplateName: params.TemplateName})
	} else if err != nil {
		log.Fatal(&tf_errors.InternalError{Cause: err, Name: tf_errors.StorageError})
	}

	// Create the root folder for the template
	err = os.MkdirAll(params.TargetDirPath, fs.FileMode(params.DirPermission))
	if err != nil {
		log.Fatal(&tf_errors.NodeCreationError{Path: params.TargetDirPath})
	}
}
