package utils

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestASCIIArtTool(t *testing.T) {
	// 1. Setup: Build the binary first to test it as a CLI tool
	build := exec.Command("go", "build", "-o", "ascii-tool")
	if err := build.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("ascii-tool") // Clean up after tests

	// 2. Define Test Cases
	tests := []struct {
		name     string
		args     []string
		contains string // What we expect to see in the output
	}{
		{
			name:     "Basic Generation",
			args:     []string{"Hello", "standard"},
			contains: "_  _  ____  _     _     __  ", // Part of 'H' and 'e' in standard
		},
		{
			name:     "Color Flag",
			args:     []string{"--color=red", "Test"},
			contains: "\x1b[31m", // ANSI escape code for red
		},
		{
			name:     "Output to File",
			args:     []string{"--output=test_out.txt", "FileTest"},
			contains: "", // Output goes to file, not console
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("./ascii-tool", tt.args...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("Command failed: %v", err)
			}

			if tt.name == "Output to File" {
				if _, err := os.Stat("test_out.txt"); os.IsNotExist(err) {
					t.Error("Output file was not created")
				}
				os.Remove("test_out.txt") // Cleanup
			} else if !strings.Contains(string(out), tt.contains) {
				t.Errorf("Expected output to contain %q, but got %q", tt.contains, string(out))
			}
		})
	}
}

func TestReverseMode(t *testing.T) {
	// Create a dummy ASCII file to test reverse
	dummyArt := " _ \n| |\n|_|\n" // Simplified example
	os.WriteFile("reverse_me.txt", []byte(dummyArt), 0644)
	defer os.Remove("reverse_me.txt")

	t.Run("Reverse Flag Check", func(t *testing.T) {
		cmd := exec.Command("go", "run", ".", "--reverse=reverse_me.txt")
		out, _ := cmd.CombinedOutput()

		// Since actual reverse logic depends on your map,
		// we check if it triggers without crashing.
		if strings.Contains(string(out), "Error") {
			t.Errorf("Reverse failed with error: %s", string(out))
		}
	})
}
