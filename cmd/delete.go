/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/delete"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete template_name",
	Short: "Delete a template",
	Long:  `Delete a particular template, if it exists`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		delete.DeleteTemplate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCommand)
}
