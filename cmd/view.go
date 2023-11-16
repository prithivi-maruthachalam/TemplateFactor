/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/view"
	"github.com/spf13/cobra"
)

var viewCommand = &cobra.Command{
	Use:   "view template_name",
	Short: "Display a template",
	Long:  `Display a particular template, if it exists`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		view.ViewTemplate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(viewCommand)
}
