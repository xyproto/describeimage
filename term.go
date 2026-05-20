package main

import (
	"os"

	"golang.org/x/term"
)

// getTerminalWidth tries to find the current width of the terminal, with a fallback to 79
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 79 // fallback value
	}
	return width
}
