package internal

import "fmt"

const SourceDirStatError = "SourceDirStatError"
const PatternMatchingError = "PatternMatchingError"
const FilePathWalkError = "FilePathWalkError"
const FileReadError = "FileReadError"

// Template name isn't valid.
type InvalidTemplateNameError struct {
	TemplateName string
}

func (err *InvalidTemplateNameError) Error() string {
	return BoldError("\nInvalidTemplateNameError : ") +
		BoldError(fmt.Sprintf("The name '%s' is Invalid", err.TemplateName)) +
		ErrorInfo("\nA valid template has to start with a letter and can only contain letters, numbers and underscores")
}

// Source directory was not found
type SourceDirNotFoundErr struct {
	SourceDir string
}

func (err *SourceDirNotFoundErr) Error() string {
	return BoldError("\nSourceDirNotFoundErr : ") +
		BoldError(fmt.Sprintf("The directory '%s' was not found", err.SourceDir))
}

// Internal error caused by some other function
type InternalError struct {
	Cause error
	Name  string
}

func (err *InternalError) Error() string {
	return BoldError("\nInternalError : ") + BoldError(err.Name) +
		ErrorInfo(fmt.Sprintf("\n%v", err.Cause))
}
