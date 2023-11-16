package io

import (
	"github.com/fatih/color"
)

var Log_Warning = color.New(color.FgYellow).SprintFunc()
var Log_Info = color.New(color.FgCyan).SprintFunc()
var Log_Debug = color.New(color.FgHiBlack).SprintFunc()
