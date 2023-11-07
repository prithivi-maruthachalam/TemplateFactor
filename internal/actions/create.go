package actions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal"
)

// Contains the input configuration used to create a template
type CreateTemplateConfig struct {
	TemplateName       string
	SourceDirPath      string `default:"."`
	SaveFiles          bool   `default:"false"`
	SaveContent        bool   `default:"false"`
	StoreLink          bool   `default:"false"`
	Clobber            bool   `default:"false"`
	DryRun             bool   `default:"false"`
	ExcludeList        []string
	FileIncludeList    []string
	ContentExcludeList []string
	ContentIncludeList []string
}

func (config *CreateTemplateConfig) String() string {
	str := fmt.Sprintf("\n  Template Name: %s\n  Source Dir: %s\n  SaveFile?: %t\n  SaveContent?: %t\n  StoreLink?: %t\n  Clobber?: %t\n  DryRun?: %t\n",
		config.TemplateName,
		config.SourceDirPath,
		config.SaveFiles,
		config.SaveContent,
		config.StoreLink,
		config.Clobber,
		config.DryRun)

	str += "  ExcludeList: ("
	for _, pat := range config.ExcludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += "  FileIncludeList: ("
	for _, pat := range config.FileIncludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += "  ContentExcludeList: ("
	for _, pat := range config.ContentExcludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += "  ContentIncludeList: ("
	for _, pat := range config.ContentIncludeList {
		str += fmt.Sprintf("'%s' ", pat)
	}
	str += ")\n"

	str += strings.Repeat("-", 10)
	str += "\n"

	return str
}

// Validates the input configuration. Returns an error if invalid.
func (config *CreateTemplateConfig) Validate() error {
	if !internal.ValidateTemplateName(config.TemplateName) {
		return &internal.InvalidTemplateNameError{TemplateName: config.TemplateName}
	}

	if _, err := os.Stat(config.SourceDirPath); os.IsNotExist(err) {
		return &internal.SourceDirNotFoundErr{SourceDir: config.SourceDirPath}
	} else if err != nil {
		return &internal.InternalError{Cause: err, Name: internal.SourceDirStatError}
	}

	return nil
}

// Creates a template from a given input configuration
func CreateTemplate(params CreateTemplateConfig) {
	fmt.Println(internal.Info("Creating Template"))

	// validate input configuration
	err := params.Validate()
	if err != nil {
		log.Fatal(err)
	}

	// setup a new template object
	newTemplate := internal.Template{
		TemplateName: params.TemplateName,
		Nodes:        []internal.TemplateNode{},
	}

	// Tests for matches in a list of glob patterns against a path
	testMatches := func(patterns []string, full_path string) bool {
		match, err := internal.TestGlobMatches(patterns, full_path)

		if err != nil {
			log.Fatal(internal.InternalError{Cause: err, Name: internal.PatternMatchingError})
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
				return &internal.InternalError{Cause: err, Name: internal.FilePathWalkError}
			}

			if path == params.SourceDirPath {
				return nil
			}

			relative_path := strings.Replace(path, params.SourceDirPath, "", 1)
			relative_path = strings.TrimPrefix(relative_path, "/")
			fmt.Println(internal.Debug(relative_path))

			if !testExclude(relative_path) {
				/* This path doesn't match any exclude pattern
				 * and can be included in the template
				 */

				if info.IsDir() {
					// If this is a directory, add it to the template
					newTemplate.AddNode(internal.TemplateNode{
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
							// File should be added to the template
							fileContent := ""

							if (params.SaveContent && !testContentExclude(relative_path)) || (!params.SaveContent && testContentInclude(relative_path)) {
								data, err := os.ReadFile(path)
								if err != nil {
									return &internal.InternalError{Cause: err, Name: internal.FileReadError}
								}

								fileContent = string(data)
							}

							newTemplate.AddNode(internal.TemplateNode{
								NodePath: relative_path,
								IsFile:   true,
								Content:  fileContent,
							})
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

	spew.Dump(params)
	spew.Dump(newTemplate)
}
