/*
Copyright © 2023 Prithivi Maruthachalam <prithivimaruthachalam@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/actions/use"
	"github.com/spf13/cobra"
)

var useCommandConstants = struct {
	TargetDirPath  FlagDef // String
	NoFiles        FlagDef // Bool
	NoFileContent  FlagDef // Bool
	DirPermission  FlagDef // Int
	FilePermission FlagDef // Int
}{
	TargetDirPath:  FlagDef{"target", "t", "The target directory that is created from the template. [Default .]", StringFlagDefault{"."}},
	NoFiles:        FlagDef{"no-file", "f", "If set, any files and their content that are a part of the template will be excluded. [Default false]", BoolFlagDefault{false}},
	NoFileContent:  FlagDef{"no-content", "F", "If set, any file content included in the template will be excluded. [Default false]", BoolFlagDefault{false}},
	DirPermission:  FlagDef{"dir-permission", "d", "The permission bits for the directories that are created using this template. [Default 0777]", IntFlagDefault{0777}},
	FilePermission: FlagDef{"file-permission", "p", "The permission bits for the files that are created using this template. [Default 0777]", IntFlagDefault{0777}},
}

var useCmd = &cobra.Command{
	Use:   "use [flags] source_dir",
	Short: "Use a template",
	Long:  `Use an existing template to create dirs and files based on the template.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params := use.UseTemplateConfig{
			TemplateName:   args[0],
			TargetDirPath:  useCommandConstants.TargetDirPath.GetStringAndHandleErr(cmd),
			NoFiles:        useCommandConstants.NoFiles.GetBoolAndHandleError(cmd),
			NoFileContent:  useCommandConstants.NoFileContent.GetBoolAndHandleError(cmd),
			DirPermission:  useCommandConstants.DirPermission.GetIntAndHandleError(cmd),
			FilePermission: useCommandConstants.FilePermission.GetIntAndHandleError(cmd),
		}

		use.UseTemplate(params)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Target dir path
	useCmd.Flags().StringP(useCommandConstants.TargetDirPath.Long,
		useCommandConstants.TargetDirPath.Short,
		useCommandConstants.TargetDirPath.Default.GetDefaultString(),
		useCommandConstants.TargetDirPath.Help)
	err := useCmd.MarkFlagDirname(useCommandConstants.TargetDirPath.Long)
	if err != nil {
		log.Fatal(err)
	}

	// NoFiles flag
	useCmd.Flags().BoolP(useCommandConstants.NoFiles.Long,
		useCommandConstants.NoFiles.Short,
		useCommandConstants.NoFiles.Default.GetDefaultBool(),
		useCommandConstants.NoFiles.Help)

	// NoContent flag
	useCmd.Flags().BoolP(useCommandConstants.NoFileContent.Long,
		useCommandConstants.NoFileContent.Short,
		useCommandConstants.NoFileContent.Default.GetDefaultBool(),
		useCommandConstants.NoFileContent.Help)

	// Dir permissions
	useCmd.Flags().IntP(useCommandConstants.DirPermission.Long,
		useCommandConstants.DirPermission.Short,
		useCommandConstants.DirPermission.Default.GetDefaultInt(),
		useCommandConstants.DirPermission.Help)

	// File permissions
	useCmd.Flags().IntP(useCommandConstants.FilePermission.Long,
		useCommandConstants.FilePermission.Short,
		useCommandConstants.FilePermission.Default.GetDefaultInt(),
		useCommandConstants.FilePermission.Help)

}
