package io

import (
	"github.com/fatih/color"
)

var BoldError = color.New(color.FgRed, color.Bold).SprintFunc()
var ErrorInfo = color.New(color.FgRed).SprintFunc()
var Warning = color.New(color.FgYellow).SprintFunc()
var Info = color.New(color.FgCyan).SprintFunc()
var Debug = color.New(color.FgHiBlack).SprintFunc()

// For displaying template
var Title = color.New(color.Bold).SprintFunc()
var SubtleText = color.New(color.FgHiBlack).SprintfFunc()
var DirName = color.New(color.Bold, color.FgCyan).SprintfFunc()
