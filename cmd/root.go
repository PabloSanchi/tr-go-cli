package cmd

import (
	"io"
	"os"
	"fmt"
	"log"
	"bufio"
	"github.com/spf13/cobra"
)

var ExecuteDeleteFlag bool
var ExecuteSqueezeFlag bool
var output string
var err error

var rootCmd = &cobra.Command{
    Use:   "cc43",
    Short: "Code Challenge 43",
    Long: "Code Challenge 43",
    Run: func(cmd *cobra.Command, args []string) {
		for {

			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {

				if err == io.EOF {
					os.Exit(0)
				}
				
				log.Fatal("An error occurred while reading input")
			}
			
			if !ExecuteDeleteFlag && !ExecuteSqueezeFlag {
				output, err = ExecuteDefault(input, args)
			}
	
			if err != nil {
				log.Fatal("Error: %w", err)
			}

			fmt.Print(output)
		}
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&ExecuteDeleteFlag, "delete", "d", false, "Execute Delete Mode")
	rootCmd.PersistentFlags().BoolVarP(&ExecuteSqueezeFlag, "squeeze", "s", false, "Execute Squeeze Mode")
}
