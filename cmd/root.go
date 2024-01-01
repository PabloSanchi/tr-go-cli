package cmd

import (
	"os"
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

var ExecuteDeleteFlag bool
var ExecuteSqueezeFlag bool
var ExecuteDeleteAndSqueezeFlag bool
var output string
var err error

var rootCmd = &cobra.Command{
    Use:   "cctr",
    Short: "Code Challenge 43",
    Long: `
	Code Challenge 43:
	NAME
		cctr – translate characters

	SYNOPSIS
		cctr [-Ccsu] string1 string2
		tcctr [-Ccu] -d string1
		cctr [-Ccu] -s string1
		cctr [-Ccu] -ds string1 string2

	DESCRIPTION
		The cctr utility copies the standard input to the standard output with
		substitution or deletion of selected characters.`,
    Run: func(cmd *cobra.Command, args []string) {

		for {
			
			input, err := ReadUserInput(args)

			if err != nil {
				log.Fatal("Error: %w", err)
			}

			output, err := ExecuteDefault(input, args)
			
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
	rootCmd.PersistentFlags().BoolVar(&ExecuteDeleteAndSqueezeFlag, "ds", false, "Execute Delete and Squeeze Mode")
}
