/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions"
	"github.com/spf13/cobra"
)

var commandConstants = struct {
	TemplateName       FlagDef // StringArray
	SaveFiles          FlagDef // Bool
	SaveContent        FlagDef // Bool
	StoreLink          FlagDef // Bool
	Clobber            FlagDef // Bool
	ExcludeList        FlagDef // StringArray
	FileIncludeList    FlagDef // StringArray
	ContentIncludeList FlagDef // StringArray
	ContentExcludeList FlagDef // StringArray
	DryRun             FlagDef // Bool
}{
	TemplateName:       FlagDef{"name", "n"},
	SaveFiles:          FlagDef{"save-files", "f"},
	SaveContent:        FlagDef{"save-content", "F"},
	StoreLink:          FlagDef{"store-link", "L"},
	Clobber:            FlagDef{"clobber", "x"},
	FileIncludeList:    FlagDef{"include-file", "i"},
	ExcludeList:        FlagDef{"exclude", "I"},
	ContentIncludeList: FlagDef{"content-include", "c"},
	ContentExcludeList: FlagDef{"content-exclude", "C"},
	DryRun:             FlagDef{"dry-run", "d"},
}

var createCmd = &cobra.Command{
	Use:   "create [flags] source_dir",
	Short: "Create a template from a source directory",
	Long:  `Create a template from a source directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params := actions.CreateTemplateConfig{
			TemplateName:    args[0],
			SourceDirPath:   GetStringAndHandleErr(cmd.Flags().GetString(commandConstants.TemplateName.Long)),
			SaveFiles:       GetBoolAndHandleError(cmd.Flags().GetBool(commandConstants.SaveFiles.Long)) || GetBoolAndHandleError(cmd.Flags().GetBool(commandConstants.SaveContent.Long)),
			SaveFileContent: GetBoolAndHandleError(cmd.Flags().GetBool(commandConstants.SaveContent.Long)),
			StoreLink:       GetBoolAndHandleError(cmd.Flags().GetBool(commandConstants.StoreLink.Long)),
			Clobber:         GetBoolAndHandleError(cmd.Flags().GetBool(commandConstants.Clobber.Long)),
			ExcludeList:     GetStringArrayAndHandleError(cmd.Flags().GetStringArray(commandConstants.ExcludeList.Long)),
		}
		actions.CreateTemplate(params)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Name of the template
	createCmd.Flags().StringP(commandConstants.TemplateName.Long,
		commandConstants.TemplateName.Short,
		"",
		"The name of the template")
	createCmd.MarkFlagRequired(commandConstants.TemplateName.Long)

	// If true, files will also be considered for the template (include and exclude patterns will apply)
	createCmd.Flags().BoolP(commandConstants.SaveFiles.Long,
		commandConstants.SaveFiles.Short,
		false,
		"If files should also be included")

	/* If true, the contents of files will also be considered for the
	 * template (include and exclude patterns will apply).
	 * Note : This has the effect of setting save_files to true
	 */
	createCmd.Flags().BoolP(commandConstants.SaveContent.Long,
		commandConstants.SaveContent.Short,
		false,
		"If the content of files should also be saved")

	/* If true, only a link to the source directory is stored. A
	 * copy of the dir is not created for the template.
	 * Note : This implies that changes to the original source dir
	 * will change the template
	 */
	createCmd.Flags().BoolP(commandConstants.StoreLink.Long,
		commandConstants.StoreLink.Short,
		false,
		"Only the links to the source are stored and no copies are created")

	// If a template with the same name exists, it will be overwritten
	createCmd.Flags().BoolP(commandConstants.Clobber.Long,
		commandConstants.Clobber.Short,
		false,
		"Force overwrite if a template with the same name exists")

	// Paths matching any of these patterns will be ignored from the template
	createCmd.Flags().StringArrayP(commandConstants.ExcludeList.Long,
		commandConstants.ExcludeList.Short,
		[]string{},
		"A list of patterns to ignore")

	// Files to include, if saveFiles and saveFileContent are false
	help_str := fmt.Sprintf("A list of glob patterns for files to include if '%s' is false", commandConstants.SaveFiles.Long)
	createCmd.Flags().StringArrayP(commandConstants.FileIncludeList.Long,
		commandConstants.FileIncludeList.Short,
		[]string{},
		help_str)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.FileIncludeList.Long, commandConstants.SaveFiles.Long)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.FileIncludeList.Long, commandConstants.SaveContent.Long)

	/* Files whose content to include if saveFileContent is false.
	 * Note: These files will be included in the template even if they
	 * are not, either by setting saveFiles to false or through EXCLUDE_LIST
	 */
	help_str = fmt.Sprintf("A list of glob patterns for file content (and original file) to include even if '%s' is false or if '%s' is false or if the file is ignored through '%s",
		commandConstants.SaveFiles.Long,
		commandConstants.SaveContent.Long,
		commandConstants.ExcludeList.Long)
	createCmd.Flags().StringArrayP(commandConstants.ContentIncludeList.Long,
		commandConstants.ContentIncludeList.Short,
		[]string{},
		help_str)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.ContentIncludeList.Long, commandConstants.SaveContent.Long)

	// List of glob patterns for files whose content is to be excluded, even if saveFileContent is true
	help_str = fmt.Sprintf("A list of glob patterns for file contents to exclude, even if '%s' is true", commandConstants.SaveContent.Long)
	createCmd.Flags().StringArrayP(commandConstants.ContentExcludeList.Long,
		commandConstants.ContentExcludeList.Short,
		[]string{},
		help_str)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.ContentExcludeList.Long, commandConstants.ContentIncludeList.Long)

	createCmd.Flags().BoolP(commandConstants.DryRun.Long,
		commandConstants.DryRun.Short,
		false,
		"Do not create a template. Only show the paths that will be included in the template")
}
