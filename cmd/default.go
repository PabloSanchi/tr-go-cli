package cmd

import (
	"fmt"
	"strings"
)


func ExecuteDefault(input string, args []string) (string, error) {
	
	if len(args) < 2 {
		return "", fmt.Errorf("At least two arguments are required")
	}

	output := strings.Replace(input, args[0], args[1], -1)

	return output, nil
}