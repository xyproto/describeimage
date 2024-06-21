package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/wordwrap"
)

// Only print the provided data when verbose is true
func logVerbose(verbose bool, format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}

// describeImages uses Ollama and the given model to describe one or more images.
// prompt is the start of the multimodal prompt: the instructions which will be followed by one or more images
// outputFile is the file to write the result to (can be blank to not write to anything)
// model is the model to use, like llava
// writeWidth is the width that the returned or written string should be wrapped to, if it is >0
// filenames is a list of input images
// A description is returned as a string.
func describeImages(prompt, outputFile, model string, wrapWidth int, filenames []string, verbose bool) (string, error) {
	if wrapWidth == -1 {
		wrapWidth = getTerminalWidth()
	}

	if len(filenames) < 1 {
		return "", fmt.Errorf("no image filenames provided")
	}

	var images []string
	for _, filename := range filenames {
		logVerbose(verbose, "[%s] Reading... ", filename)
		base64image, err := ollamaclient.Base64EncodeFile(filename)
		if err == nil { // success
			images = append(images, base64image)
			logVerbose(verbose, "OK\n")
		} else {
			logVerbose(verbose, "FAILED: "+err.Error()+"\n")
		}
	}

	if len(images) == 0 {
		return "", fmt.Errorf("no images to describe")
	}

	if prompt == "" {
		if len(images) > 1 {
			prompt = "Describe the following images:"
		} else {
			prompt = "Describe this image:"
		}
	}

	oc := ollamaclient.New()
	oc.ModelName = model
	oc.Verbose = verbose

	if err := oc.PullIfNeeded(verbose); err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	oc.SetReproducible()

	promptAndImages := append([]string{prompt}, images...)

	logVerbose(verbose, "[%s] Analyzing...\n", oc.ModelName)
	output, err := oc.GetOutput(promptAndImages...)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	logVerbose(verbose, "[%s] Analysis complete.\n", oc.ModelName)

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
