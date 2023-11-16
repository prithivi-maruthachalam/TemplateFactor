/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/view"
	"github.com/spf13/cobra"
)

var ShowContentFlag = FlagDef{"show-content", "s", "If set, the content for files will also be displayed."}

var viewCommand = &cobra.Command{
	Use:   "view [flags] template_name",
	Short: "Display a template",
	Long:  `Display a particular template, if it exists`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		view.ViewTemplate(args[0], ShowContentFlag.GetBoolAndHandleError(cmd))
	},
}

func init() {
	rootCmd.AddCommand(viewCommand)

	viewCommand.Flags().BoolP(ShowContentFlag.Long,
		ShowContentFlag.Short,
		false,
		ShowContentFlag.Help)
}
