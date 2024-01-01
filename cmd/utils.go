package cmd

import (
	"os"
	"io"
	"fmt"
	"bufio"
)

func ReadUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var input string

	input, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			os.Exit(0)
		}
		return "", fmt.Errorf("An error occurred while reading input: %v", err)
	}


	return input, nil
}
