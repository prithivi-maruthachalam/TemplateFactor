package errors

import (
	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Path to os home was not found
type HomePathNotFoundError struct {
}

func (err *HomePathNotFoundError) Error() string {
	return tf_io.Error_Title("HomePathNotFoundError : ") +
		tf_io.Error_Info("Couldn't fine home path for your system. Check that the env variable for Home is set correctly for your operating system")
}
