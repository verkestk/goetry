package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/verkestk/markovokram"

	"github.com/verkestk/goetry/src/corpus"
	"github.com/verkestk/goetry/src/util/markov"
)

var sentencePerson string
var sentenceLength int

var generateSentencesCmd = &cobra.Command{
	Use:   "generate-sentences",
	Short: "generates sentences",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cor, _, err := corpus.Load(corpusFilepath, sentencePerson)
		if err != nil {
			return fmt.Errorf("error loading corpus: %w", err)
		}

		rand.Seed(time.Now().UnixNano())
		chain := markovokram.NewChain(prefixLength)
		for _, line := range cor.Lines {
			chain.Build(strings.Fields(line))
		}
		text := markov.GenerateSentences(chain, sentenceLength)
		fmt.Println(text)
		return nil
	},
}

func init() {
	generateSentencesCmd.Flags().StringVarP(&corpusFilepath, "corpus", "c", "", "path to the corpus file")
	generateSentencesCmd.Flags().StringVarP(&sentencePerson, "person", "p", "", "person to base the generated text from")
	generateSentencesCmd.Flags().IntVarP(&sentenceLength, "length", "l", 1, "number of sentences to generate")
	generateSentencesCmd.Flags().IntVarP(&prefixLength, "prefix-length", "", 2, "length of markov chain prefix")
	generateSentencesCmd.MarkFlagRequired("corpus")
	rootCmd.AddCommand(generateSentencesCmd)
}
