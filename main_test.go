package main

import (
	"strings"
	"testing"
	"time"
)

const (
	promptHeader = ""
	outputFile   = ""
	wrapWidth    = 0
	verbose      = true
)

func TestImageDescriptionContainsHorses(t *testing.T) {
	// Define the input parameters
	filenames := []string{"img/horses.png"}
	// Call the describeImages function
	output, err := describeImages("", "", "", wrapWidth, filenames, verbose)
	if err != nil {
		t.Fatalf("describeImages failed: %v", err)
	}
	// Check if the output contains the word "horses"
	if !strings.Contains(output, "horses") {
		t.Errorf("Expected output to contain 'horses', but it did not. Output: %s", output)
	}
}

func TestImageDescriptionContainsCat(t *testing.T) {
	// Define the input parameters
	filenames := []string{"img/meloncat.jpg"}
	// Call the describeImages function
	time.Sleep(10 * time.Second)
	output, err := describeImages("", "", "", wrapWidth, filenames, verbose)
	if err != nil {
		t.Fatalf("describeImages failed: %v", err)
	}
	// Check if the output contains the word "cat"
	if !strings.Contains(output, "cat") {
		t.Errorf("Expected output to contain 'cat', but it did not. Output: %s", output)
	}
}
