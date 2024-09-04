package main

import (
	"fmt"
	"strings"
)

func Encoder(line string) string {
	lines := strings.Split(line, "\n")
	encodedLines := make([]string, len(lines))
	for i, line := range lines {
		encodedLines[i] = encodeLine(line)
	}
	return strings.Join(encodedLines, "\n")
}

func encodeLine(line string) string {
	line = removeNonPrintables(line)
	if line == "" {
		return ""
	}
	encoded := ""
	count := 1
	for i := 1; i < len(line); i++ {
		if line[i] == line[i-1] {
			count++
		} else {
			if count > 1 {
				encoded += fmt.Sprintf("[%d %c]", count, line[i-1])
			} else {
				encoded += string(line[i-1])
			}
			count = 1
		}
	}
	if count > 1 {
		encoded += fmt.Sprintf("[%d %c]", count, line[len(line)-1])
	} else {
		encoded += string(line[len(line)-1])
	}
	return encoded
}
