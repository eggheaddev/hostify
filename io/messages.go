package io

import (
	"fmt"
)

// Reset red ANSI color
const Reset = "\u001b[0m"

// Red ANSI color
const Red = "\u001b[31m"

// Green ANSI color
const Green = "\u001b[32m"

// Yellow ANSI color
const Yellow = "\u001b[33m"

var bug = Yellow + "\nlooks like a bug? if you think this, please send us your report:" + Reset + Green + "\n==> https://github.com/eggheaddev/hostify-cli/issues\n" + Reset

// Trace default message
var Trace = bug + Red + "\n ==== Error trace ==== \n" + Reset

// ErrorMessage print error message
func ErrorMessage(text string) {
	fmt.Println(Red + "Error: " + Reset + text)
}

// SuccessMessage print success message
func SuccessMessage(text string) {
	fmt.Println(Green + "Done: " + Reset + text)
}

// WarningMessage print warning message
func WarningMessage(text string) {
	fmt.Println(Yellow + "Warning: " + Reset + text)
}
