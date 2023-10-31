/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions"
	"github.com/spf13/cobra"
)

type CommandDef struct {
	Long  string
	Short string
}

var commandConstants = struct {
	TEMPLATE_NAME        CommandDef // StringArray
	SAVE_FILES           CommandDef // Bool
	SAVE_CONTENT         CommandDef // Bool
	STORE_LINK           CommandDef // Bool
	CLOBBER              CommandDef // Bool
	EXCLUDE_LIST         CommandDef // StringArray
	INCLUDE_LIST         CommandDef // StringArray
	INCLUDE_CONTENT_LIST CommandDef // StringArray
	EXCLUDE_CONTENT_LIST CommandDef // StringArray
	DRY_RUN              CommandDef // Bool
}{
	TEMPLATE_NAME:        CommandDef{"name", "n"},
	SAVE_FILES:           CommandDef{"save_files", "f"},
	SAVE_CONTENT:         CommandDef{"save_content", "F"},
	STORE_LINK:           CommandDef{"store_link", "L"},
	CLOBBER:              CommandDef{"clobber", "x"},
	INCLUDE_LIST:         CommandDef{"include_files", "i"},
	EXCLUDE_LIST:         CommandDef{"exclude", "I"},
	INCLUDE_CONTENT_LIST: CommandDef{"include_content", "c"},
	EXCLUDE_CONTENT_LIST: CommandDef{"exclude_content", "C"},
	DRY_RUN:              CommandDef{"dry_run", "d"},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [flags] source_dir",
	Short: "Create a template from a source directory",
	Long:  `Create a template from a source directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src, _ := cmd.Flags().GetString(commandConstants.TEMPLATE_NAME.Long)
		sf, _ := cmd.Flags().GetBool(commandConstants.SAVE_FILES.Long)
		sfc, _ := cmd.Flags().GetBool(commandConstants.SAVE_CONTENT.Long)
		sl, _ := cmd.Flags().GetBool(commandConstants.STORE_LINK.Long)
		cl, _ := cmd.Flags().GetBool(commandConstants.CLOBBER.Long)
		EXCLUDE_LIST, _ := cmd.Flags().GetStringArray(commandConstants.EXCLUDE_LIST.Long)

		params := actions.CreateTemplateConfig{
			TemplateName:    args[0],
			SourceDirPath:   src,
			SaveFiles:       sf || sfc,
			SaveFileContent: sfc,
			StoreLinks:      sl,
			Clobber:         cl,
			IgnoreList:      EXCLUDE_LIST,
		}
		actions.CreateTemplate(params)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Name of the template
	createCmd.Flags().StringP(commandConstants.TEMPLATE_NAME.Long,
		commandConstants.TEMPLATE_NAME.Short,
		"",
		"The name of the template")
	createCmd.MarkFlagRequired(commandConstants.TEMPLATE_NAME.Long)

	// If true, files will also be considered for the template (include and exclude patterns will apply)
	createCmd.Flags().BoolP(commandConstants.SAVE_FILES.Long,
		commandConstants.SAVE_FILES.Short,
		false,
		"If files should also be included")

	/* If true, the contents of files will also be considered for the
	 * template (include and exclude patterns will apply).
	 * Note : This has the effect of setting save_files to true
	 */
	createCmd.Flags().BoolP(commandConstants.SAVE_CONTENT.Long,
		commandConstants.SAVE_CONTENT.Short,
		false,
		"If the content of files should also be saved")

	/* If true, only a link to the source directory is stored. A
	 * copy of the dir is not created for the template.
	 * Note : This implies that changes to the original source dir
	 * will change the template
	 */
	createCmd.Flags().BoolP(commandConstants.STORE_LINK.Long,
		commandConstants.STORE_LINK.Short,
		false,
		"Only the links to the source are stored and no copies are created")

	// If a template with the same name exists, it will be overwritten
	createCmd.Flags().BoolP(commandConstants.CLOBBER.Long,
		commandConstants.CLOBBER.Short,
		false,
		"Force overwrite if a template with the same name exists")

	// Paths matching any of these patterns will be ignored from the template
	createCmd.Flags().StringArrayP(commandConstants.EXCLUDE_LIST.Long,
		commandConstants.EXCLUDE_LIST.Short,
		[]string{},
		"A list of patterns to ignore")

	// Files to include, if saveFiles and saveFileContent are false
	help_str := fmt.Sprintf("A list of glob patterns for files to include if '%s' is false", commandConstants.SAVE_FILES.Long)
	createCmd.Flags().StringArrayP(commandConstants.INCLUDE_LIST.Long,
		commandConstants.INCLUDE_LIST.Short,
		[]string{},
		help_str)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.INCLUDE_LIST.Long, commandConstants.SAVE_FILES.Long)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.INCLUDE_LIST.Long, commandConstants.SAVE_CONTENT.Long)

	/* Files whose content to include if saveFileContent is false.
	 * Note: These files will be included in the template even if they
	 * are not, either by setting saveFiles to false or through EXCLUDE_LIST
	 */
	help_str = fmt.Sprintf("A list of glob patterns for file content (and original file) to include even if '%s' is false or if '%s' is false or if the file is ignored through '%s",
		commandConstants.SAVE_FILES.Long,
		commandConstants.SAVE_CONTENT.Long,
		commandConstants.EXCLUDE_LIST.Long)
	createCmd.Flags().StringArrayP(commandConstants.INCLUDE_CONTENT_LIST.Long,
		commandConstants.INCLUDE_CONTENT_LIST.Short,
		[]string{},
		help_str)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.INCLUDE_CONTENT_LIST.Long, commandConstants.SAVE_CONTENT.Long)

	// List of glob patterns for files whose content is to be excluded, even if saveFileContent is true
	help_str = fmt.Sprintf("A list of glob patterns for file contents to exclude, even if '%s' is true", commandConstants.SAVE_CONTENT.Long)
	createCmd.Flags().StringArrayP(commandConstants.EXCLUDE_CONTENT_LIST.Long,
		commandConstants.EXCLUDE_CONTENT_LIST.Short,
		[]string{},
		help_str)
	createCmd.MarkFlagsMutuallyExclusive(commandConstants.EXCLUDE_CONTENT_LIST.Long, commandConstants.INCLUDE_CONTENT_LIST.Long)

	createCmd.Flags().BoolP(commandConstants.DRY_RUN.Long,
		commandConstants.DRY_RUN.Short,
		false,
		"Do not create a template. Only show the paths that will be included in the template")
}
