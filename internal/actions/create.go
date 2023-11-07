package actions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal"
)

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

func testGlobMatches(patterns []string, path string) (bool, error) {
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

func CreateTemplate(params CreateTemplateConfig) {
	color.Blue("Creating Template")

	// Validate Template Name
	if !internal.ValidateTemplateName(params.TemplateName) {
		log.Fatal(&internal.InvalidTemplateNameError{TemplateName: params.TemplateName})
	}

	// Validate that source dir exists
	if _, err := os.Stat(params.SourceDirPath); os.IsNotExist(err) {
		log.Fatal(err)
	}

	newTemplate := internal.Template{
		TemplateName: params.TemplateName,
		Nodes:        []internal.TemplateNode{},
	}

	testExclude := func(full_path string) bool {
		match, err := testGlobMatches(params.ExcludeList, full_path)

		if err != nil {
			log.Fatal(err)
		}

		return match
	}

	testFileInclude := func(full_path string) bool {
		match, err := testGlobMatches(params.FileIncludeList, full_path)

		if err != nil {
			log.Fatal(err)
		}

		return match
	}

	testContentInclude := func(full_path string) bool {
		match, err := testGlobMatches(params.ContentIncludeList, full_path)

		if err != nil {
			log.Fatal(err)
		}

		return match
	}

	testContentExclude := func(full_path string) bool {
		match, err := testGlobMatches(params.ContentExcludeList, full_path)

		if err != nil {
			log.Fatal(err)
		}

		return match
	}

	// Recursively go through every file and dir in the source directory
	err := filepath.Walk(params.SourceDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if path == params.SourceDirPath {
				return nil
			}

			relative_path := strings.Replace(path, params.SourceDirPath, "", 1)
			relative_path = strings.TrimPrefix(relative_path, "/")
			fmt.Println(relative_path)

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
								fileContent = "including content"
							}

							newTemplate.AddNode(internal.TemplateNode{
								NodePath: relative_path,
								IsFile:   true,
								Content:  fileContent,
							})
						}
					}
				}
			} else {
				if info.IsDir() {
					return filepath.SkipDir
				}

				return nil
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\nDeubgging")
	spew.Dump(params)
	spew.Dump(newTemplate)
}
