package errors

import (
	"fmt"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Source directory was not found
type SourceDirNotFoundError struct {
	SourceDir string
}

func (err *SourceDirNotFoundError) Error() string {
	return tf_io.Error_Title("SourceDirNotFoundError : ") +
		tf_io.Error_Title(fmt.Sprintf("The directory '%s' was not found", err.SourceDir))
}
