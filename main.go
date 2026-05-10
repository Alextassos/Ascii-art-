package main

import (
	"fmt"
	"os"
	"reverse/utils"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	cfg := utils.Config{Banner: "standard", Align: "left"}
	var words []string

	// Ενιαίο Loop για όλα τα Flags
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--reverse=") {
			cfg.ReverseFile = strings.TrimPrefix(arg, "--reverse=")
		} else if strings.HasPrefix(arg, "--color=") {
			cfg.Color = strings.TrimPrefix(arg, "--color=")
		} else if strings.HasPrefix(arg, "--align=") {
			cfg.Align = strings.TrimPrefix(arg, "--align=")
			if cfg.Align != "left" && cfg.Align != "right" && cfg.Align != "center" && cfg.Align != "justify" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
				return
			}
		} else if strings.HasPrefix(arg, "--output=") {
			cfg.OutputFile = strings.TrimPrefix(arg, "--output=")
		} else {
			words = append(words, arg)
		}
	}

	// ΠΕΡΙΠΤΩΣΗ 1: REVERSE
	if cfg.ReverseFile != "" {
		// Αν το flag είναι σκέτο "--reverse=", βγάλε usage
		if cfg.ReverseFile == "" {
			fmt.Println("Usage: go run . --reverse=<fileName>")
			return
		}
		result, err := utils.HandleReverse(cfg.ReverseFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(result)
		return
	}

	// ΠΕΡΙΠΤΩΣΗ 2: ΚΑΝΟΝΙΚΟ ASCII ART
	if len(words) == 0 {
		return
	}

	// Διαχωρισμός Text/Banner/Substr
	if cfg.Color != "" && len(words) >= 2 {
		cfg.Substr = words[0]
		cfg.Text = words[1]
		if len(words) == 3 {
			cfg.Banner = words[2]
		}
	} else {
		cfg.Text = words[0]
		if len(words) >= 2 {
			cfg.Banner = words[1]
		}
	}

	// Validations
	if !utils.IsValidBanner(cfg.Banner) {
		fmt.Println("Error: invalid banner")
		return
	}

	// Διάβασμα Banner
	bannerPath := cfg.Banner + ".txt"
	content, err := os.ReadFile(bannerPath)
	if err != nil {
		fmt.Println("Error: Banner not found")
		return
	}

	bannerLines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	rawText := strings.ReplaceAll(cfg.Text, "\\n", "\n")
	textLines := strings.Split(rawText, "\n")

	var finalOutput strings.Builder
	for _, line := range textLines {
		colorMap := utils.GetColorMap(line, cfg.Substr)
		finalOutput.WriteString(utils.BuildStyledLine(line, bannerLines, cfg, colorMap))
	}

	// Output (File or Console)
	if cfg.OutputFile != "" {
		err = os.WriteFile(cfg.OutputFile, []byte(finalOutput.String()), 0644)
		if err != nil {
			fmt.Println("Error writing file")
		}
	} else {
		fmt.Print(finalOutput.String())
	}

}
