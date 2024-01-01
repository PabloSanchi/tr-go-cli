package cmd

import (
	"fmt"
	"strings"
	"golang.org/x/exp/slices"
)

func MapChar(char string, from, to []string) string {
	idx := slices.IndexFunc(from, func(s string) bool { return s == char })

	if idx == -1 {
		return char
	}

	if idx >= len(to) {
		return to[len(to) - 1]
	}

	return to[idx]
}

func TranslateChars(from, to []string, input string) string {
	var output strings.Builder
	for _, c := range input {
		output.WriteString(MapChar(string(c), from, to))
	}
	return output.String()
}

func ToSlice(s string) ([]string, error) {
	var slice []string
	for _, r := range s {
		slice = append(slice, string(r))
	}
	return slice, nil
}

func ToRangeSlice(arg string) ([]string, error) {
	firstChar, lastChar := rune(arg[0]), rune(arg[len(arg)-1])

	if lastChar < firstChar {
		return []string{}, fmt.Errorf("Invalid range")
	}

	var slice []string
	for i := firstChar; i <= lastChar; i++ {
		slice = append(slice, string(i))
	}

	return slice, nil
}

func ExecuteDefault(input string, args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("At least two arguments are required")
	}

	firstArg, err := parseArg(args[0])

	if err != nil {
		return "", fmt.Errorf("First argument parse error: %w", err)
	}

	secondArg, err := parseArg(args[1])

	if err != nil {
		return "", fmt.Errorf("Second argument parse error: %w", err)
	}

	return TranslateChars(firstArg, secondArg, input), nil
}

func parseArg(arg string) ([]string, error) {
	if strings.Contains(arg, "-") {
		return ToRangeSlice(arg)
	}
	return ToSlice(arg)
}
