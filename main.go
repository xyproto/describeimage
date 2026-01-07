package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const versionString = "DescribeImage 1.3.3"

func main() {
	var (
		outputFile   string
		promptHeader string
		wrapWidth    int
		showVersion  bool
		verbose      bool
		model        string
	)

	pflag.BoolVarP(&verbose, "verbose", "V", false, "verbose output")
	pflag.StringVarP(&model, "model", "m", "", "Specify the Ollama model to use")
	pflag.StringVarP(&promptHeader, "prompt", "p", "Describe the following image(s):", "Provide a custom prompt header")
	pflag.StringVarP(&outputFile, "output", "o", "", "Specify an output file")
	pflag.IntVarP(&wrapWidth, "wrap", "w", 0, "Word wrap at specified width. Use '-1' for terminal width")
	pflag.BoolVarP(&showVersion, "version", "v", false, "display version")

	pflag.Parse()

	if showVersion {
		fmt.Println(versionString)
		return
	}

	filenames := pflag.Args()

	output, err := describeImages(promptHeader, model, outputFile, wrapWidth, filenames, verbose)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if output != "" {
		fmt.Println(output)
	}
}
