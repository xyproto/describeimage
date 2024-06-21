package main

import (
	"strings"
	"testing"
)

const (
	promptHeader = ""
	outputFile   = ""
	wrapWidth    = 0
)

func TestImageDescriptionContainsPuppy(t *testing.T) {
	// Define the input parameters
	filenames := []string{"img/puppy.png"}
	// Call the describeImages function
	output, err := describeImages(promptHeader, outputFile, defaultModel, wrapWidth, filenames)
	if err != nil {
		t.Fatalf("describeImages failed: %v", err)
	}
	// Check if the output contains the word "puppy"
	if !strings.Contains(output, "puppy") {
		t.Errorf("Expected output to contain 'puppy', but it did not. Output: %s", output)
	}
}

func TestImageDescriptionContainsCat(t *testing.T) {
	// Define the input parameters
	filenames := []string{"img/meloncat.jpg"}
	// Call the describeImages function
	output, err := describeImages(promptHeader, outputFile, defaultModel, wrapWidth, filenames)
	if err != nil {
		t.Fatalf("describeImages failed: %v", err)
	}
	// Check if the output contains the word "cat"
	if !strings.Contains(output, "cat") {
		t.Errorf("Expected output to contain 'cat', but it did not. Output: %s", output)
	}
}
