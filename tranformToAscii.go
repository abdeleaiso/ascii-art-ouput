package main

import (
	"strings"
)

// Transform input into ascii art using the banner font in lines.
func TransformToAscii(input string, lines []string) string {
	if input == "" {
		return ""
	}

	// The shell passes "\n" as two characters (\ and n), so turn it
	// Into a real newline before splitting
	input = strings.ReplaceAll(input, "\\n", "\n")
	inputLines := strings.Split(input, "\n")

	// If the input is only "\n\n...", there is nothing to draw.
	// Just print that many blanks lines.
	isAllEmpty := true
	for i := 0; i < len(inputLines); i++ {
		if inputLines[i] != "" {
			isAllEmpty = false
			break
		}
	}
	if isAllEmpty {
		return strings.Repeat("\n", len(inputLines)-1)
	}

	var b strings.Builder

	for _, line := range inputLines {
		// Empty line in the input becomes one empty line in the output.
		if line == "" {
			b.WriteByte('\n')
			continue
		}
		// Go through every char and append its piece of that row.
		for i := 0; i < 8; i++ {
			for _, c := range line {
				// Only printable ASCII.
				if c < 32 || c > 126 {
					continue
				}
				// 9 lines per char , +1 to skip the blank separator line.
				startIndex := (int(c)-32)*9 + 1
				// Safety check so short/corrupted font file doesn't panic.
				if startIndex+i < len(lines) {
					b.WriteString(lines[startIndex+i])
				}
			}
			b.WriteByte('\n')
		}
	}
	return b.String()
}
