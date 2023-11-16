/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/storage"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfac",
	Short: "Create reusable, reconfigurable templates out of your directories and files!",
	Long: `TemplateFactory allows you to create 'templates' out of your directories, files and their contents. You have the option to create templates with just dir structures, just empty files or even include the content of a few files. You can then use your templates to recreate those directories and files anywhere else!

Remember that there are no limits to what can be part of a template. Use the include and exclude patterns to fully customize what goes into your templates!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.templatefactory.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	err := storage.TestAndCreateTemplateFactoryHome()
	if err != nil {
		log.Fatal(&errors.TemplateFactoryHomeCreationError{Path: storage.TF_HOME})
	}
}
