package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// Only print the provided data when verbose is true
func logVerbose(verbose bool, format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}

// getTerminalWidth tries to find the current width of the terminal, with a fallback to 79
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 79 // fallback value
	}
	return width
}
