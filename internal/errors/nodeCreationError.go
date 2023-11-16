package errors

import (
	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Path to os home was not found
type NodeCreationError struct {
	Path string
}

func (err *NodeCreationError) Error() string {
	return tf_io.Error_Title("NodeCreationError : ") +
		tf_io.Error_Info("Error creating item '%s'", err.Path)
}
