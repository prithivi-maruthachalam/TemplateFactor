/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions"
	"github.com/spf13/cobra"
)

const TEMPLATE_NAME = "name"
const SAVE_FILES = "save_files"
const SAVE_CONTENT = "save_content"
const STORE_LINK = "store_link"
const FORCE = "force"
const CONFIG_FILE = "config_file"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [flags] source_dir",
	Short: "Create a template from a source directory",
	Long:  `Create a template from a source directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src, _ := cmd.Flags().GetString(TEMPLATE_NAME)
		sf, _ := cmd.Flags().GetBool(SAVE_FILES)
		sfc, _ := cmd.Flags().GetBool(SAVE_CONTENT)
		sl, _ := cmd.Flags().GetBool(STORE_LINK)
		force, _ := cmd.Flags().GetBool(FORCE)
		cp, _ := cmd.Flags().GetString(CONFIG_FILE)

		params := actions.CreateTemplateConfig{
			TemplateName:    args[0],
			SourceDirPath:   src,
			SaveFiles:       sf,
			SaveFileContent: sfc,
			StoreLinks:      sl,
			Force:           force,
			ConfigPath:      cp,
		}
		actions.CreateTemplate(params)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP(TEMPLATE_NAME, "n", "", "The name of the template")
	createCmd.MarkFlagRequired(TEMPLATE_NAME)

	createCmd.Flags().BoolP(SAVE_FILES, "f", false, "If files should also be included")
	createCmd.Flags().BoolP(SAVE_CONTENT, "c", false, "If the content of files should also be saved")
	createCmd.Flags().BoolP(STORE_LINK, "l", false, "Only the links to the source are stored and no copies are created")
	createCmd.Flags().BoolP(FORCE, "x", false, "Force overwrite if a template with the same name exists")
	createCmd.Flags().StringP(CONFIG_FILE, "e", "", "Path to the config file to use for includes/excludes")
}
