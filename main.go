package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/wordwrap"
	"golang.org/x/term"
)

const (
	versionString    = "DescribeImage 1.0.1"
	defaultModel     = "llava"
	defaultTermWidth = 79
)

var verbose bool

// Only print the provided data when in verbose mode
func logVerbose(format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return defaultTermWidth
	}
	return width
}

func describeImages(promptHeader, outputFile, model string, wrapWidth int, filenames []string) (string, error) {
	if wrapWidth == -1 {
		wrapWidth = getTerminalWidth()
	}

	if len(filenames) < 1 {
		return "", fmt.Errorf("no image filenames provided")
	}

	var images []string
	for _, filename := range filenames {
		logVerbose("[%s] Reading... ", filename)
		base64image, err := ollamaclient.Base64EncodeFile(filename)
		if err == nil { // success
			images = append(images, base64image)
			logVerbose("OK\n")
		} else {
			logVerbose("FAILED: " + err.Error() + "\n")
		}
	}

	var prompt string
	switch len(images) {
	case 0:
		return "", fmt.Errorf("no images to describe")
	case 1:
		prompt = "Describe this image:"
	default:
		prompt = "Describe these images:"
	}
	if promptHeader != "" {
		prompt = promptHeader
	}

	oc := ollamaclient.New()
	oc.ModelName = model

	if err := oc.PullIfNeeded(verbose); err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	oc.SetReproducible()

	promptAndImages := append([]string{prompt}, images...)

	logVerbose("[%s] Generating... ", oc.ModelName)
	output, err := oc.GetOutput(promptAndImages...)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	logVerbose("OK\n")

	if output == "" {
		return "", fmt.Errorf("generated output for the prompt %s is empty", prompt)
	}

	if wrapWidth > 0 {
		lines, err := wordwrap.WordWrap(output, wrapWidth)
		if err == nil { // success
			output = strings.Join(lines, "\n")
		}
	}

	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(output), 0o644)
		if err != nil {
			return "", fmt.Errorf("error writing to file: %v", err)
		}
		return "", nil
	}

	return output, nil
}

func main() {
	var (
		promptHeader, outputFile, model string
		wrapWidth                       int
		showVersion                     bool
	)

	pflag.BoolVarP(&verbose, "verbose", "V", false, "verbose output")
	pflag.StringVarP(&promptHeader, "prompt", "p", "Write a short summary of what a project that contains the following files is:", "Provide a custom prompt header")
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

	output, err := describeImages(promptHeader, outputFile, model, wrapWidth, filenames)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if output != "" {
		fmt.Println(output)
	}
}
