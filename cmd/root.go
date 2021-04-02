package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goetry",
	Short: "Goetry generates poems using a hidden markov model.",
	Long:  "Goetry is a CLI gor generating different sorts of poetry based on a corpus of real text, using a markov model. Feed it a corpus, a pronuciaion dictionary, and some poetic form definitions, and see what it can do!\n\nhttps://github.com/verkestk/goetry",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to geotry!\n Run with `--help` for instructions.")
	},
}

// Execute executes a CLI command - boilerplate for cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
