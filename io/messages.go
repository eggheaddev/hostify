package io

import (
	"fmt"
)

const reset = "\u001b[0m"
const red = "\u001b[31m"
const green = "\u001b[32m"

// ErrorMessage print error message
func ErrorMessage(text string) {
	fmt.Println(red + "Error: " + reset + text)
}

// SuccessMessage print success message
func SuccessMessage(text string) {
	fmt.Println(green + "Done: " + reset + text)
}
