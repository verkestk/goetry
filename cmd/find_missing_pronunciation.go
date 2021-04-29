package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/verkestk/goetry/src/corpus"
	"github.com/verkestk/goetry/src/rhymes"
)

var findMissingPronunciationCmd = &cobra.Command{
	Use:   "find-missing-pronunciation",
	Short: "reports all words from corpus that have no pronunication in the dictionary",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cor, _, err := corpus.Load(corpusFilepath, "")
		if err != nil {
			return fmt.Errorf("error loading corpus: %w", err)
		}

		rhymer, err := rhymes.Load(pronunciationDictionaryFilepath, cor)
		if err != nil {
			return fmt.Errorf("error loading rhymer: %w", err)
		}

		missingPronunciation := rhymer.UnknownPronunciations()
		if len(missingPronunciation) == 0 {
			fmt.Println("There are no unknown pronunciations in the corpus.")
		} else {
			fmt.Println("Pronunciation missing for the following words:")
			for _, word := range missingPronunciation {
				fmt.Printf("  %s\n", word)
			}
		}

		return nil
	},
}

func init() {
	findMissingPronunciationCmd.Flags().StringVarP(&corpusFilepath, "corpus", "c", "", "path to the corpus file")
	findMissingPronunciationCmd.Flags().StringVarP(&pronunciationDictionaryFilepath, "dictionary", "d", "", "path to the pronunciation dictionary file")
	findMissingPronunciationCmd.MarkFlagRequired("corpus")
	findMissingPronunciationCmd.MarkFlagRequired("dictionary")
	rootCmd.AddCommand(findMissingPronunciationCmd)
}
