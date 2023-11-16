package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
)

// Validates the name of the template
func ValidateTemplateName(name string) bool {
	match, err := regexp.MatchString("^[a-zA-Z$][a-zA-Z_$0-9]+$", name)

	if err != nil {
		log.Fatal(err)
	}

	return match
}

// Test if the given path string matches any of the glob patterns in the slice
func TestGlobMatches(patterns []string, path string) (bool, error) {
	for _, pattern := range patterns {
		match, err := doublestar.Match(pattern, path)
		if err != nil {
			return false, err
		}

		if match {
			return match, nil
		}
	}

	return false, nil
}

// Gets the home directory for the given operating system
func GetUserHomeDir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(errors.HomePathNotFound{})
	}

	return homedir
}
