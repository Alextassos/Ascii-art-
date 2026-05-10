package utils

import (
	"fmt"
	"os"
	"strings"
)

func HandleReverse(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// 1. Καθαρίζουμε τα line endings και αφαιρούμε τις κενές γραμμές
	// ΜΟΝΟ από την αρχή και το τέλος του αρχείου.
	rawStr := strings.ReplaceAll(string(content), "\r\n", "\n")
	trimmedStr := strings.Trim(rawStr, "\n")
	lines := strings.Split(trimmedStr, "\n")

	// Αν το αρχείο είναι άδειο μετά το trim
	if len(lines) < 8 {
		return "", fmt.Errorf("file too short or empty")
	}

	banners := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}

	for _, bName := range banners {
		bannerContent, err := os.ReadFile(bName)
		if err != nil {
			continue
		}

		bannerLines := strings.Split(strings.ReplaceAll(string(bannerContent), "\r\n", "\n"), "\n")

		// map: pattern (8 lines string) -> rune
		patterns := make(map[string]rune)

		for i := 0; i < 95; i++ {
			var pattern strings.Builder
			for j := 1; j <= 8; j++ {
				idx := i*9 + j
				if idx < len(bannerLines) {
					pattern.WriteString(bannerLines[idx])
					pattern.WriteString("\n")
				}
			}
			patterns[pattern.String()] = rune(i + 32)
		}

		var result strings.Builder
		// Επεξεργασία ανά μπλοκ 8 γραμμών
		for row := 0; row <= len(lines)-8; row += 8 {

			// Βρίσκουμε το μέγιστο μήκος γραμμής στο τρέχον μπλοκ
			maxLen := 0
			for i := 0; i < 8; i++ {
				if len(lines[row+i]) > maxLen {
					maxLen = len(lines[row+i])
				}
			}

			for col := 0; col < maxLen; {
				found := false

				// Δοκιμάζουμε πλάτη (widths) για να βρούμε match
				// Το 20 είναι ένα ασφαλές max πλάτος για shadow/standard γράμματα
				for width := 1; width <= 30 && col+width <= maxLen; width++ {
					var candidate strings.Builder

					for i := 0; i < 8; i++ {
						line := lines[row+i]

						// Παίρνουμε το segment και προσθέτουμε padding αν η γραμμή είναι μικρότερη
						segment := ""
						if col < len(line) {
							end := col + width
							if end > len(line) {
								end = len(line)
							}
							segment = line[col:end]
						}

						if len(segment) < width {
							segment += strings.Repeat(" ", width-len(segment))
						}

						candidate.WriteString(segment)
						candidate.WriteString("\n")
					}

					if ch, ok := patterns[candidate.String()]; ok {
						result.WriteRune(ch)
						col += width
						found = true
						break
					}
				}

				if !found {
					col++ // Αν δεν βρει τίποτα, προχωράει ένα κενό
				}
			}
			result.WriteString("\n")
		}

		if result.Len() > 0 {
			// Επιστρέφουμε το αποτέλεσμα καθαρό
			finalOutput := strings.TrimSpace(result.String())
			if finalOutput != "" {
				return finalOutput, nil
			}
		}
	}

	return "", fmt.Errorf("no match found")
}
