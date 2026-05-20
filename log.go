package main

import (
	"fmt"
)

// Only print the provided data when verbose is true
func logVerbose(verbose bool, format string, a ...any) {
	if verbose {
		fmt.Printf(format, a...)
	}
}
