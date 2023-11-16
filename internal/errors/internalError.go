package errors

import (
	"fmt"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

const SourceDirStatError = "SourceDirStatError"
const PatternMatchingError = "PatternMatchingError"
const FilePathWalkError = "FilePathWalkError"
const FileReadError = "FileReadError"

// Internal error caused by some other function
type InternalError struct {
	Cause error
	Name  string
}

func (err *InternalError) Error() string {
	return tf_io.Error_Title("\nInternalError : ") + tf_io.Error_Title(err.Name) +
		tf_io.Error_Info(fmt.Sprintf("\n%v", err.Cause))
}
