package errors

import (
	"fmt"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Source directory was not found
type TemplateNotFound struct {
	TemplateName string
}

func (err *TemplateNotFound) Error() string {
	return tf_io.Error_Title("TemplateNotFound : ") +
		tf_io.Error_Info(fmt.Sprintf("The template '%s' was not found", err.TemplateName))
}
