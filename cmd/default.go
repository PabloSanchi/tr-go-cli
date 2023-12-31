package cmd

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func mapCharacter(c string, from []string, to []string) string {
	idx := getIndex(from, c)

	if idx == -1 {
		return c
	}

	if idx >= len(to) {
		return to[len(to) - 1]
	}

	return to[idx]
}

func getIndex(slice []string, target string) int {
	idx := slices.IndexFunc(slice, func(s string) bool { return s == target })
	return idx
}

func TRCharacters(from []string, to []string, input string) string {
	var output string
	for _, c := range input {
		output += mapCharacter(string(c), from, to)
	}
	return output
}

func DefaultStrategy(arg string) []string {
	var charSlice []string
    for _, r := range arg {
        charSlice = append(charSlice, string(r))
    }
	return charSlice
}

/*
stategies:
	default => no - or "[:<class>:]"
	range -> that means the arument contains a -
	class -> that means the arument contains a "[:<class>:]"
*/
func ExecuteDefault(input string, args []string) (string, error) {
	
	if len(args) < 2 {
		return "", fmt.Errorf("At least two arguments are required")
	}

	firstArg := DefaultStrategy(args[0])
	secondArg := DefaultStrategy(args[1])

	fmt.Println("firstArg", firstArg)
	fmt.Println("secondArg", secondArg)

	output := TRCharacters(firstArg, secondArg, input)

	return output, nil
}