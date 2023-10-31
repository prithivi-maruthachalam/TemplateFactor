package cmd

import "log"

type FlagDef struct {
	Long  string
	Short string
}

func GetStringAndHandleErr(v string, err error) string {
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func GetBoolAndHandleError(v bool, err error) bool {
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func GetStringArrayAndHandleError(v []string, err error) []string {
	if err != nil {
		log.Fatal(err)
	}

	return v
}
