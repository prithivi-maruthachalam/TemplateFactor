package errors

import (
	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Error creating template factory home directory
type TemplateFactoryHomeCreationError struct {
	Path string
}

func (err *TemplateFactoryHomeCreationError) Error() string {
	return tf_io.Error_Title("TemplateFactoryHomeCreationError : ") +
		tf_io.Error_Info("Error creating template factory home directory at '%s'", err.Path)
}
