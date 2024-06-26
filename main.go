package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"golang.org/x/term"
)

const (
	versionString    = "DescribeImage 1.2.0"
	defaultModel     = "llava"
	defaultTermWidth = 79
)

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return defaultTermWidth
	}
	return width
}

func main() {
	var (
		promptHeader, outputFile, model string
		wrapWidth                       int
		showVersion                     bool
		verbose                         bool
	)

	pflag.BoolVarP(&verbose, "verbose", "V", false, "verbose output")
	pflag.StringVarP(&promptHeader, "prompt", "p", "Describe the following image(s):", "Provide a custom prompt header")
	pflag.StringVarP(&outputFile, "output", "o", "", "Specify an output file")
	pflag.StringVarP(&model, "model", "m", defaultModel, "Specify the Ollama model to use")
	pflag.IntVarP(&wrapWidth, "wrap", "w", 0, "Word wrap at specified width. Use '-1' for terminal width")
	pflag.BoolVarP(&showVersion, "version", "v", false, "display version")

	pflag.Parse()

	if showVersion {
		fmt.Println(versionString)
		return
	}

	filenames := pflag.Args()

	output, err := describeImages(promptHeader, outputFile, model, wrapWidth, filenames, verbose)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if output != "" {
		fmt.Println(output)
	}
}
