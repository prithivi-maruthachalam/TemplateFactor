/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/create"
	"github.com/spf13/cobra"
)

var commandConstants = struct {
	TemplateName       FlagDef // StringArray
	SaveFiles          FlagDef // Bool
	SaveContent        FlagDef // Bool
	Clobber            FlagDef // Bool
	ExcludeList        FlagDef // StringArray
	FileIncludeList    FlagDef // StringArray
	ContentIncludeList FlagDef // StringArray
	ContentExcludeList FlagDef // StringArray
	DryRun             FlagDef // Bool
}{
	TemplateName: FlagDef{"name", "n", "The name of the template"},
	SaveFiles:    FlagDef{"save-files", "f", "If set, files will also be saved to the template. Only directories are part of the template by default. [Default false]"},
	SaveContent:  FlagDef{"save-content", "F", "If set, files and their content will be included in the template. [Default false]"},
	Clobber:      FlagDef{"clobber", "x", "If set, an existing template with the same name will be overwritten without warning. [Default false]"},
	FileIncludeList: FlagDef{"include-file", "i", `A list of glob patterns for files that should be included in the template, even if save-files is false.
Can't be used with save-files or save-content`},
	ExcludeList: FlagDef{"exclude", "e", "A set of glob patterns for directories and files to be excluded from the template. This overrides all other include/exclude options."},
	ContentIncludeList: FlagDef{"content-include", "c", `A list of glob patterns for files whose content will be included in the template, even if save-files, or save-content are false.
Can't be used with content-exclude or with save-content`},
	ContentExcludeList: FlagDef{"content-exclude", "C", `A list of glob patterns for files whose content will be excluded from the template, even if save-content is true.
Can't be used with content-include.`},
	DryRun: FlagDef{"dry-run", "d", "If set, a template will be displayed, but not created. [Default false]"},
}

var createCmd = &cobra.Command{
	Use:   "create [flags] source_dir",
	Short: "Create a template from a source directory",
	Long:  `Create a template from a source directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params := create.CreateTemplateConfig{
			SourceDirPath:      args[0],
			TemplateName:       commandConstants.TemplateName.GetStringAndHandleErr(cmd),
			SaveFiles:          commandConstants.SaveFiles.GetBoolAndHandleError(cmd),
			SaveContent:        commandConstants.SaveContent.GetBoolAndHandleError(cmd),
			Clobber:            commandConstants.Clobber.GetBoolAndHandleError(cmd),
			DryRun:             commandConstants.DryRun.GetBoolAndHandleError(cmd),
			ExcludeList:        commandConstants.ExcludeList.GetStringArrayAndHandleError(cmd),
			FileIncludeList:    commandConstants.FileIncludeList.GetStringArrayAndHandleError(cmd),
			ContentExcludeList: commandConstants.ContentExcludeList.GetStringArrayAndHandleError(cmd),
			ContentIncludeList: commandConstants.ContentIncludeList.GetStringArrayAndHandleError(cmd),
		}
		create.CreateTemplate(params)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Name of the template
	createCmd.Flags().StringP(commandConstants.TemplateName.Long,
		commandConstants.TemplateName.Short,
		"",
		commandConstants.TemplateName.Help)
	if err := createCmd.MarkFlagRequired(commandConstants.TemplateName.Long); err != nil {
		log.Fatal(err)
	}

	// If true, files will also be considered for the template (include and exclude patterns will apply)
	createCmd.Flags().BoolP(commandConstants.SaveFiles.Long,
		commandConstants.SaveFiles.Short,
		false,
		commandConstants.SaveFiles.Help)

	/* If true, the contents of files will also be considered for the
	 * template (include and exclude patterns will apply).
	 * Note : This has the effect of setting save_files to true
	 */
	createCmd.Flags().BoolP(commandConstants.SaveContent.Long,
		commandConstants.SaveContent.Short,
		false,
		commandConstants.SaveContent.Help)

	// If a template with the same name exists, it will be overwritten
	createCmd.Flags().BoolP(commandConstants.Clobber.Long,
		commandConstants.Clobber.Short,
		false,
		commandConstants.Clobber.Help)

	// Paths matching any of these patterns will be ignored from the template
	createCmd.Flags().StringArrayP(commandConstants.ExcludeList.Long,
		commandConstants.ExcludeList.Short,
		[]string{},
		commandConstants.ExcludeList.Help)

	// Files to include, if saveFiles and saveFileContent are false
	createCmd.Flags().StringArrayP(commandConstants.FileIncludeList.Long,
		commandConstants.FileIncludeList.Short,
		[]string{},
		commandConstants.FileIncludeList.Help)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.FileIncludeList.Long, commandConstants.SaveFiles.Long)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.FileIncludeList.Long, commandConstants.SaveContent.Long)

	/* Files whose content to include if saveFileContent is false.
	 * Note: These files will be included in the template even if they
	 * are not, either by setting saveFiles to false or through EXCLUDE_LIST
	 */
	createCmd.Flags().StringArrayP(commandConstants.ContentIncludeList.Long,
		commandConstants.ContentIncludeList.Short,
		[]string{},
		commandConstants.ContentIncludeList.Help)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.ContentIncludeList.Long, commandConstants.SaveContent.Long)

	// List of glob patterns for files whose content is to be excluded, even if saveFileContent is true
	createCmd.Flags().StringArrayP(commandConstants.ContentExcludeList.Long,
		commandConstants.ContentExcludeList.Short,
		[]string{},
		commandConstants.ContentExcludeList.Help)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.ContentExcludeList.Long, commandConstants.ContentIncludeList.Long)

	createCmd.Flags().BoolP(commandConstants.DryRun.Long,
		commandConstants.DryRun.Short,
		false,
		commandConstants.DryRun.Help)
}
