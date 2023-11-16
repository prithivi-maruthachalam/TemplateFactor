package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type FlagDefault interface {
	GetDefaultInt() int
	GetDefaultString() string
	GetDefaultBool() bool
	GetDefaultStringArray() []string
}

type FlagDef struct {
	Long    string
	Short   string
	Help    string
	Default FlagDefault
}

type IntFlagDefault struct {
	value int
}

func (def IntFlagDefault) GetDefaultInt() int {
	return def.value
}
func (def IntFlagDefault) GetDefaultString() string {
	return ""
}
func (def IntFlagDefault) GetDefaultBool() bool {
	return false
}
func (def IntFlagDefault) GetDefaultStringArray() []string {
	return nil
}

type StringFlagDefault struct {
	value string
}

func (def StringFlagDefault) GetDefaultInt() int {
	return 0
}
func (def StringFlagDefault) GetDefaultString() string {
	return def.value
}
func (def StringFlagDefault) GetDefaultBool() bool {
	return false
}
func (def StringFlagDefault) GetDefaultStringArray() []string {
	return nil
}

type BoolFlagDefault struct {
	value bool
}

func (def BoolFlagDefault) GetDefaultInt() int {
	return 0
}
func (def BoolFlagDefault) GetDefaultString() string {
	return ""
}
func (def BoolFlagDefault) GetDefaultBool() bool {
	return def.value
}
func (def BoolFlagDefault) GetDefaultStringArray() []string {
	return nil
}

type StringArrayFlagDefault struct {
	value []string
}

func (def StringArrayFlagDefault) GetDefaultInt() int {
	return 0
}
func (def StringArrayFlagDefault) GetDefaultString() string {
	return ""
}
func (def StringArrayFlagDefault) GetDefaultBool() bool {
	return false
}
func (def StringArrayFlagDefault) GetDefaultStringArray() []string {
	return def.value
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

func (flag *FlagDef) GetIntAndHandleError(cmd *cobra.Command) int {
	v, err := cmd.Flags().GetInt(flag.Long)
	if err != nil {
		log.Fatal(err)
	}

	return v
}
