package io

import (
	"fmt"
)

// Reset red ANSI color
const Reset = "\u001b[0m"

// Red red ANSI color
const Red = "\u001b[31m"

// Green red ANSI color
const Green = "\u001b[32m"

// ErrorMessage print error message
func ErrorMessage(text string) {
	fmt.Println(Red + "Error: " + Reset + text)
}

// SuccessMessage print success message
func SuccessMessage(text string) {
	fmt.Println(Green + "Done: " + Reset + text)
}
