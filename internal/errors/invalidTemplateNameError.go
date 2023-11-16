package errors

import (
	"fmt"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Template name isn't valid.
type InvalidTemplateNameError struct {
	TemplateName string
}

func (err *InvalidTemplateNameError) Error() string {
	return tf_io.Error_Title("\nInvalidTemplateNameError : ") +
		tf_io.Error_Title(fmt.Sprintf("The name '%s' is Invalid", err.TemplateName)) +
		tf_io.Error_Info("\nA valid template has to start with a letter and can only contain letters, numbers and underscores")
}
