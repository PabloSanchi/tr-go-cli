package cmd

import (
	"os"
	"io"
	"fmt"
	"bufio"
)

func ReadUserInput(args []string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {

		if err == io.EOF {
			os.Exit(0)
		}
		
		return "", fmt.Errorf("An error occurred while reading input")
	}

	return input, nil	
}