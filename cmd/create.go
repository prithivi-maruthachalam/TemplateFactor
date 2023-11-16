/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/create"
	"github.com/spf13/cobra"
)

var createCommandConstants = struct {
	SourceDirPath      FlagDef // String
	SaveFiles          FlagDef // Bool
	SaveContent        FlagDef // Bool
	Clobber            FlagDef // Bool
	ExcludeList        FlagDef // StringArray
	FileIncludeList    FlagDef // StringArray
	ContentIncludeList FlagDef // StringArray
	ContentExcludeList FlagDef // StringArray
	DryRun             FlagDef // Bool
}{
	SourceDirPath: FlagDef{"source", "s", "The source directory to build the template from. [Default .]", StringFlagDefault{"."}},
	SaveFiles:     FlagDef{"save-files", "f", "If set, files will also be saved to the template. Only directories are part of the template by default. [Default false]", BoolFlagDefault{false}},
	SaveContent:   FlagDef{"save-content", "F", "If set, files and their content will be included in the template. [Default false]", BoolFlagDefault{false}},
	Clobber:       FlagDef{"clobber", "x", "If set, an existing template with the same name will be overwritten without warning. [Default false]", BoolFlagDefault{false}},
	FileIncludeList: FlagDef{"include-file", "i", `A list of glob patterns for files that should be included in the template, even if save-files is false.
Can't be used with save-files or save-content`, StringArrayFlagDefault{[]string{}}},
	ExcludeList: FlagDef{"exclude", "e", "A set of glob patterns for directories and files to be excluded from the template. This overrides all other include/exclude options.", StringArrayFlagDefault{[]string{}}},
	ContentIncludeList: FlagDef{"content-include", "c", `A list of glob patterns for files whose content will be included in the template, even if save-files, or save-content are false.
Can't be used with content-exclude or with save-content`, StringArrayFlagDefault{[]string{}}},
	ContentExcludeList: FlagDef{"content-exclude", "C", `A list of glob patterns for files whose content will be excluded from the template, even if save-content is true.
Can't be used with content-include.`, StringArrayFlagDefault{[]string{}}},
	DryRun: FlagDef{"dry-run", "d", "If set, a template will be displayed, but not created. [Default false]", BoolFlagDefault{false}},
}

var createCmd = &cobra.Command{
	Use:   "create [flags] template_name",
	Short: "Create a template from a source directory",
	Long:  `Create a template from a source directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params := create.CreateTemplateConfig{
			SourceDirPath:      createCommandConstants.SourceDirPath.GetStringAndHandleErr(cmd),
			TemplateName:       args[0],
			SaveFiles:          createCommandConstants.SaveFiles.GetBoolAndHandleError(cmd),
			SaveContent:        createCommandConstants.SaveContent.GetBoolAndHandleError(cmd),
			Clobber:            createCommandConstants.Clobber.GetBoolAndHandleError(cmd),
			DryRun:             createCommandConstants.DryRun.GetBoolAndHandleError(cmd),
			ExcludeList:        createCommandConstants.ExcludeList.GetStringArrayAndHandleError(cmd),
			FileIncludeList:    createCommandConstants.FileIncludeList.GetStringArrayAndHandleError(cmd),
			ContentExcludeList: createCommandConstants.ContentExcludeList.GetStringArrayAndHandleError(cmd),
			ContentIncludeList: createCommandConstants.ContentIncludeList.GetStringArrayAndHandleError(cmd),
		}
		create.CreateTemplate(params)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Source dir path
	createCmd.Flags().StringP(createCommandConstants.SourceDirPath.Long,
		createCommandConstants.SourceDirPath.Short,
		createCommandConstants.SourceDirPath.Default.GetDefaultString(),
		createCommandConstants.SourceDirPath.Help)
	err := createCmd.MarkFlagDirname(createCommandConstants.SourceDirPath.Long)
	if err != nil {
		log.Fatal(err)
	}

	// If true, files will also be considered for the template (include and exclude patterns will apply)
	createCmd.Flags().BoolP(createCommandConstants.SaveFiles.Long,
		createCommandConstants.SaveFiles.Short,
		createCommandConstants.SaveFiles.Default.GetDefaultBool(),
		createCommandConstants.SaveFiles.Help)

	/* If true, the contents of files will also be considered for the
	 * template (include and exclude patterns will apply).
	 * Note : This has the effect of setting save_files to true
	 */
	createCmd.Flags().BoolP(createCommandConstants.SaveContent.Long,
		createCommandConstants.SaveContent.Short,
		createCommandConstants.SaveContent.Default.GetDefaultBool(),
		createCommandConstants.SaveContent.Help)

	// If a template with the same name exists, it will be overwritten
	createCmd.Flags().BoolP(createCommandConstants.Clobber.Long,
		createCommandConstants.Clobber.Short,
		createCommandConstants.Clobber.Default.GetDefaultBool(),
		createCommandConstants.Clobber.Help)

	// Paths matching any of these patterns will be ignored from the template
	createCmd.Flags().StringArrayP(createCommandConstants.ExcludeList.Long,
		createCommandConstants.ExcludeList.Short,
		createCommandConstants.ExcludeList.Default.GetDefaultStringArray(),
		createCommandConstants.ExcludeList.Help)

	// Files to include, if saveFiles and saveFileContent are false
	createCmd.Flags().StringArrayP(createCommandConstants.FileIncludeList.Long,
		createCommandConstants.FileIncludeList.Short,
		createCommandConstants.FileIncludeList.Default.GetDefaultStringArray(),
		createCommandConstants.FileIncludeList.Help)
	createCmd.MarkFlagsMutuallyExclusive(createCommandConstants.FileIncludeList.Long, createCommandConstants.SaveFiles.Long)
	createCmd.MarkFlagsMutuallyExclusive(createCommandConstants.FileIncludeList.Long, createCommandConstants.SaveContent.Long)

	/* Files whose content to include if saveFileContent is false.
	 * Note: These files will be included in the template even if they
	 * are not, either by setting saveFiles to false or through EXCLUDE_LIST
	 */
	createCmd.Flags().StringArrayP(createCommandConstants.ContentIncludeList.Long,
		createCommandConstants.ContentIncludeList.Short,
		createCommandConstants.ContentIncludeList.Default.GetDefaultStringArray(),
		createCommandConstants.ContentIncludeList.Help)
	createCmd.MarkFlagsMutuallyExclusive(createCommandConstants.ContentIncludeList.Long, createCommandConstants.SaveContent.Long)

	// List of glob patterns for files whose content is to be excluded, even if saveFileContent is true
	createCmd.Flags().StringArrayP(createCommandConstants.ContentExcludeList.Long,
		createCommandConstants.ContentExcludeList.Short,
		createCommandConstants.ContentExcludeList.Default.GetDefaultStringArray(),
		createCommandConstants.ContentExcludeList.Help)
	createCmd.MarkFlagsMutuallyExclusive(createCommandConstants.ContentExcludeList.Long, createCommandConstants.ContentIncludeList.Long)

	createCmd.Flags().BoolP(createCommandConstants.DryRun.Long,
		createCommandConstants.DryRun.Short,
		createCommandConstants.DryRun.Default.GetDefaultBool(),
		createCommandConstants.DryRun.Help)
}
