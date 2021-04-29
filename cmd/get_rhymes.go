package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/verkestk/goetry/src/corpus"
	"github.com/verkestk/goetry/src/rhymes"
)

var rhymesWord string
var rhymesStrength int
var rhymesMax int

var getRhymesCmd = &cobra.Command{
	Use:   "get-rhymes",
	Short: "gets rhymes from a corpus for a word, ordered by descending strength",
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

		pronunciations := rhymer.Pronunciations(rhymesWord)
		if len(pronunciations) == 0 {
			return fmt.Errorf("rhyming dictionary missing pronuncation for %s", rhymesWord)
		}

		for _, pronunciation := range pronunciations {

			rhymes := rhymer.Rhymes(rhymesWord, pronunciation, rhymesStrength)
			if len(rhymes) > rhymesMax {
				rhymes = rhymes[:rhymesMax]
			}

			fmt.Printf("\nrhymes for %s (%s):\n", rhymesWord, strings.Join(pronunciation, " "))
			for _, rhyme := range rhymes {
				fmt.Printf("  %s (%s)\n", rhyme.Word, strings.Join(rhyme.Pronunciation, " "))
			}
		}

		return nil
	},
}

func init() {
	getRhymesCmd.Flags().StringVarP(&corpusFilepath, "corpus", "c", "", "path to the corpus file")
	getRhymesCmd.Flags().StringVarP(&pronunciationDictionaryFilepath, "dictionary", "d", "", "path to the pronunciation dictionary file")
	getRhymesCmd.Flags().StringVarP(&rhymesWord, "word", "w", "", "the word for which to find rhymes")
	getRhymesCmd.Flags().IntVarP(&rhymesStrength, "strength", "s", 1, "the minimum rhyme strength")
	getRhymesCmd.Flags().IntVarP(&rhymesMax, "max", "m", 20, "the minimum rhyme strength")
	getRhymesCmd.MarkFlagRequired("corpus")
	getRhymesCmd.MarkFlagRequired("dictionary")
	getRhymesCmd.MarkFlagRequired("word")
	rootCmd.AddCommand(getRhymesCmd)
}
