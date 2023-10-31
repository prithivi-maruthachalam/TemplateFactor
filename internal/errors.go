package internal

import "fmt"

type InvalidTemplateNameError struct {
	TemplateName string
}

func (err *InvalidTemplateNameError) Error() string {
	return fmt.Sprintf("\nInvalid Template Name\n  The name '%s is Invalid. A valid template name :\n  - Can only contain letters, numbers and underscores.\n  - Cannot start with a number or underscore\n", err.TemplateName)
}
