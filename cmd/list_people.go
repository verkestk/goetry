package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/verkestk/goetry/src/corpus"
)

var listCorpus string

var listPeopleCmd = &cobra.Command{
	Use:   "list-people",
	Short: "generates a list of people available for the generate-text command",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		_, people, err := corpus.Load(listCorpus, "")
		if err != nil {
			return fmt.Errorf("error loading corpus: %w", err)
		}

		for _, person := range people {
			fmt.Println(person)
		}

		return nil
	},
}

func init() {
	listPeopleCmd.Flags().StringVarP(&listCorpus, "corpus", "c", "", "path to the corpus file")
	listPeopleCmd.MarkFlagRequired("corpus")
	rootCmd.AddCommand(listPeopleCmd)
}
