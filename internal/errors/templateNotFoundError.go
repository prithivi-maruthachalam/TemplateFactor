package errors

import (
	"fmt"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Source directory was not found
type TemplateNotFoundError struct {
	TemplateName string
}

func (err *TemplateNotFoundError) Error() string {
	return tf_io.Error_Title("TemplateNotFoundError : ") +
		tf_io.Error_Info(fmt.Sprintf("The template '%s' was not found", err.TemplateName))
}
