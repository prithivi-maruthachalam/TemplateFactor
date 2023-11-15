package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	tf_common "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common"
	tf_errors "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common/errors"
	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common/io"
	tf_utils "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common/utils"
)

// Creates a template from a given input configuration
func CreateTemplate(params CreateTemplateConfig) {
	if params.DryRun {
		fmt.Println(tf_io.Info("Showing Template"))
	} else {
		fmt.Println(tf_io.Info("Creating Template"))
	}

	// validate input configuration
	err := params.Validate()
	if err != nil {
		log.Fatal(err)
	}

	// setup a new template object
	newTemplate := tf_common.Template{
		TemplateName: params.TemplateName,
		Nodes:        []tf_common.TemplateNode{},
	}

	// Tests for matches in a list of glob patterns against a path
	testMatches := func(patterns []string, full_path string) bool {
		match, err := tf_utils.TestGlobMatches(patterns, full_path)

		if err != nil {
			log.Fatal(tf_errors.InternalError{Cause: err, Name: tf_errors.PatternMatchingError})
		}

		return match
	}

	// tests for matches in the exclude list
	testExclude := func(full_path string) bool {
		return testMatches(params.ExcludeList, full_path)
	}

	// tests for matches in the file include list
	testFileInclude := func(full_path string) bool {
		return testMatches(params.FileIncludeList, full_path)
	}

	// tests for matches in the content include list
	testContentInclude := func(full_path string) bool {
		return testMatches(params.ContentIncludeList, full_path)
	}

	// tests for matches in the content exclude list
	testContentExclude := func(full_path string) bool {
		return testMatches(params.ContentExcludeList, full_path)
	}

	// Recursively go through every file and dir in the source directory
	err = filepath.Walk(params.SourceDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return &tf_errors.InternalError{Cause: err, Name: tf_errors.FilePathWalkError}
			}

			if path == params.SourceDirPath {
				return nil
			}

			relative_path := strings.Replace(path, params.SourceDirPath, "", 1)
			relative_path = strings.TrimPrefix(relative_path, "/")
			// fmt.Println(tf_io.Debug(relative_path))

			if !testExclude(relative_path) {
				/* This path doesn't match any exclude pattern
				 * and can be included in the template
				 */

				if info.IsDir() {
					// If this is a directory, add it to the template
					newTemplate.AddNode(tf_common.TemplateNode{
						NodePath: relative_path,
						IsFile:   false,
					})

				} else {
					/* This is a file. It needs to be added based on
					 * other parameters
					 */

					if params.SaveFiles || params.SaveContent {
						/* If saveFiles or saveContent is set to true, the file
						 * can be added to the template
						 */

						if (params.SaveFiles || params.SaveContent) || (testFileInclude(relative_path)) {
							newNode := tf_common.TemplateNode{
								NodePath:  relative_path,
								IsFile:    true,
								IsContent: false,
								Content:   "",
							}

							if (params.SaveContent && !testContentExclude(relative_path)) || (!params.SaveContent && testContentInclude(relative_path)) {
								data, err := os.ReadFile(path)
								if err != nil {
									return &tf_errors.InternalError{Cause: err, Name: tf_errors.FileReadError}
								}

								newNode.Content = string(data)
								newNode.IsContent = true

							}
							newTemplate.AddNode(newNode)
						}
					}
				}
			} else if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	// Show the template that is going to be created
	fmt.Println()
	fmt.Println(newTemplate.Describe(params.StoreLink))

	if params.DryRun {
		// If this is a dry run, return here
		return
	}
}
