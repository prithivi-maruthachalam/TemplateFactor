package errors

import (
	"fmt"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common/io"
)

const SourceDirStatError = "SourceDirStatError"
const PatternMatchingError = "PatternMatchingError"
const FilePathWalkError = "FilePathWalkError"
const FileReadError = "FileReadError"

// Template name isn't valid.
type InvalidTemplateNameError struct {
	TemplateName string
}

func (err *InvalidTemplateNameError) Error() string {
	return tf_io.BoldError("\nInvalidTemplateNameError : ") +
		tf_io.BoldError(fmt.Sprintf("The name '%s' is Invalid", err.TemplateName)) +
		tf_io.ErrorInfo("\nA valid template has to start with a letter and can only contain letters, numbers and underscores")
}

// Source directory was not found
type SourceDirNotFoundErr struct {
	SourceDir string
}

func (err *SourceDirNotFoundErr) Error() string {
	return tf_io.BoldError("\nSourceDirNotFoundErr : ") +
		tf_io.BoldError(fmt.Sprintf("The directory '%s' was not found", err.SourceDir))
}

// Internal error caused by some other function
type InternalError struct {
	Cause error
	Name  string
}

func (err *InternalError) Error() string {
	return tf_io.BoldError("\nInternalError : ") + tf_io.BoldError(err.Name) +
		tf_io.ErrorInfo(fmt.Sprintf("\n%v", err.Cause))
}
