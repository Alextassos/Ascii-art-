package utils

import (
	"strings"
)

type Config struct {
	Text, Substr, Color, Align, OutputFile, Banner string
	ReverseFile                                    string
}

func BuildStyledLine(line string, banner []string, cfg Config, colorMap []bool) string {

	if line == "" {
		return "\n"
	}

	termWidth := GetTerminalWidth()

	var finalStr strings.Builder
	words := strings.Split(line, " ")

	totalWidth := 0

	for _, char := range line {
		totalWidth += len(banner[(int(char)-32)*9+1])
	}

	padding := 0
	justifySpace := 0

	if cfg.Align == "right" {

		padding = termWidth - totalWidth

		if padding < 0 {
			padding = 0
		}

	} else if cfg.Align == "center" {

		padding = (termWidth - totalWidth) / 2

		if padding < 0 {
			padding = 0
		}

	} else if cfg.Align == "justify" && len(words) > 1 {

		remaining := termWidth - totalWidth

		if remaining < 0 {
			remaining = 0
		}

		justifySpace = remaining / (len(words) - 1)
	}

	for i := 1; i <= 8; i++ {

		finalStr.WriteString(strings.Repeat(" ", padding))

		charIdx := 0

		for wordIdx, word := range words {

			for _, char := range word {

				asciiPart := banner[(int(char)-32)*9+i]

				if cfg.OutputFile == "" && colorMap[charIdx] && cfg.Color != "" {

					finalStr.WriteString(
						GetColorCode(cfg.Color) + asciiPart + "\033[0m",
					)

				} else {

					finalStr.WriteString(asciiPart)
				}

				charIdx++
			}

			if wordIdx < len(words)-1 {

				if cfg.Align == "justify" {

					finalStr.WriteString(strings.Repeat(" ", justifySpace))

				} else {

					finalStr.WriteString(banner[(int(' ')-32)*9+i])
				}

				charIdx++
			}
		}

		finalStr.WriteString("\n")
	}

	return finalStr.String()
}
