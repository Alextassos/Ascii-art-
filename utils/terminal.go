package utils

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {

	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()

	if err != nil {
		return 80
	}

	width, err := strconv.Atoi(strings.TrimSpace(string(out)))

	if err != nil {
		return 80
	}

	return width
}

func GetColorCode(name string) string {

	colors := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"orange":  "\033[38;5;208m",
	}

	if val, ok := colors[strings.ToLower(name)]; ok {
		return val
	}

	return ""
}

func GetColorMap(text, substr string) []bool {

	isColored := make([]bool, len(text))

	if substr == "" {

		for i := range isColored {
			isColored[i] = true
		}

		return isColored
	}

	start := 0

	for {

		i := strings.Index(text[start:], substr)

		if i == -1 {
			break
		}

		pos := start + i

		for j := 0; j < len(substr); j++ {
			isColored[pos+j] = true
		}

		start = pos + 1
	}

	return isColored
}
