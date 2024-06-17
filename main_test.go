package main

import (
	"strings"
	"testing"
)

func TestImageDescriptionContainsPuppy(t *testing.T) {
	// Define the input parameters
	promptHeader := ""
	outputFile := ""
	model := defaultModel
	wrapWidth := 0
	filenames := []string{"img/puppy.png"}

	// Call the describeImages function
	output, err := describeImages(promptHeader, outputFile, model, wrapWidth, filenames)
	if err != nil {
		t.Fatalf("describeImages failed: %v", err)
	}

	// Check if the output contains the word "puppy"
	if !strings.Contains(output, "puppy") {
		t.Errorf("Expected output to contain 'puppy', but it did not. Output: %s", output)
	}
}
