package actions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/davecgh/go-spew/spew"
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
	fmt.Println("Creating Template")

	// Validate Template Name
	if !internal.ValidateTemplateName(params.TemplateName) {
		log.Fatal(&internal.InvalidTemplateNameError{TemplateName: params.TemplateName})
	}

	// Validate that source dir exists
	if _, err := os.Stat(params.SourceDirPath); os.IsNotExist(err) {
		log.Fatal(err)
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
			match, err := testGlobMatches(params.ExcludeList, relative_path)
			if err != nil {
				log.Fatal(err)
			}

			if !match {
				fmt.Println(relative_path)
			} else if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\nDeubgging")
	spew.Dump(params)
	// spew.Dump(newTemplate)
}
