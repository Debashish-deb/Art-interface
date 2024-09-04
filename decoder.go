package main

import (
	"fmt"
	"strconv"
)

func Decoder(line string) (string, error) {
	decoded := ""
	i := 0
	for i < len(line) {
		if line[i] == '[' {
			count, j, char, err := charCount(line, i)
			if err != nil {
				return "", err
			}
			for k := 0; k < count; k++ {
				decoded += char
			}
			i = j + 1
		} else if line[i] == ']' {
			return "", fmt.Errorf("unbalanced closing bracket at position %d", i)
		} else {
			decoded += string(line[i])
			i++
		}
	}
	return decoded, nil
}

func charCount(line string, index int) (int, int, string, error) {
	count := ""
	i := index + 1

	if line[index] != '[' {
		return 0, 0, "", fmt.Errorf("missing opening bracket at position %d", index)
	}
	for i < len(line) && line[i] >= '0' && line[i] <= '9' {
		count += string(line[i])
		i++
	}
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid number format")
	}
	if i >= len(line) || line[i] != ' ' {
		return 0, 0, "", fmt.Errorf("missing space between count and character")
	}
	i++
	arg := ""
	openBrackets := 1
	for i < len(line) {
		if line[i] == '[' {
			openBrackets++
		} else if line[i] == ']' {
			openBrackets--
			if openBrackets == 0 {
				break
			}
		}
		arg += string(line[i])
		i++
	}
	if openBrackets != 0 {
		return 0, 0, "", fmt.Errorf("unbalanced brackets in argument")
	}
	if arg == "" {
		return 0, 0, "", fmt.Errorf("missing character argument")
	}
	if i >= len(line) || line[i] != ']' {
		return 0, 0, "", fmt.Errorf("missing closing bracket for argument")
	}
	return countInt, i, arg, nil
}
