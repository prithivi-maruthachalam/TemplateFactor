package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type FlagDef struct {
	Long  string
	Short string
	Help  string
}

func (flag *FlagDef) GetStringAndHandleErr(cmd *cobra.Command) string {
	v, err := cmd.Flags().GetString(flag.Long)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func (flag *FlagDef) GetBoolAndHandleError(cmd *cobra.Command) bool {
	v, err := cmd.Flags().GetBool(flag.Long)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func (flag *FlagDef) GetStringArrayAndHandleError(cmd *cobra.Command) []string {
	v, err := cmd.Flags().GetStringArray(flag.Long)
	if err != nil {
		log.Fatal(err)
	}

	return v
}
