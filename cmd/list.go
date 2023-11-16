/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/list"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all templates",
	Long:  `List all templates stored in the system`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		list.ListTemplate()
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
