package internal

import (
	"github.com/fatih/color"
)

var BoldError = color.New(color.FgRed, color.Bold).SprintFunc()
var ErrorInfo = color.New(color.FgRed).SprintFunc()
var Warning = color.New(color.FgYellow).SprintFunc()
var Info = color.New(color.FgCyan).SprintFunc()
var Debug = color.New(color.FgHiBlack).SprintFunc()
