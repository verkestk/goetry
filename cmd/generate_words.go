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

var wordPerson string
var wordLength int

var generateWordsCmd = &cobra.Command{
	Use:   "generate-words",
	Short: "generates words",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cor, _, err := corpus.Load(corpusFilepath, wordPerson)
		if err != nil {
			return fmt.Errorf("error loading corpus: %w", err)
		}

		rand.Seed(time.Now().UnixNano())
		chain := markovokram.NewChain(prefixLength)
		for _, line := range cor.Lines {
			chain.Build(strings.Fields(line))
		}
		text := markov.GenerateWords(chain, wordLength)
		fmt.Println(text)
		return nil
	},
}

func init() {
	generateWordsCmd.Flags().StringVarP(&corpusFilepath, "corpus", "c", "", "path to the corpus file")
	generateWordsCmd.Flags().StringVarP(&wordPerson, "person", "p", "", "person to base the generated text from")
	generateWordsCmd.Flags().IntVarP(&wordLength, "length", "l", 10, "number of words to generate")
	generateWordsCmd.Flags().IntVarP(&prefixLength, "prefix-length", "", 2, "length of markov chain prefix")
	generateWordsCmd.MarkFlagRequired("corpus")
	rootCmd.AddCommand(generateWordsCmd)
}
