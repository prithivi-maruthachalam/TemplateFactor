package io

import "github.com/fatih/color"

// For displaying template
var Title = color.New(color.Bold).SprintFunc()
var SubtleText = color.New(color.FgHiBlack).SprintfFunc()
var DirName = color.New(color.Bold, color.FgCyan).SprintfFunc()
