package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	"github.com/verkestk/goetry/src/corpus"
	"github.com/verkestk/goetry/src/markov"
)

var wordCorpus string
var wordPerson string
var wordLength int

var generateWordsCmd = &cobra.Command{
	Use:   "generate-words",
	Short: "generates words",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cor, _, err := corpus.Load("/Users/karlieverkest/Workspace/corpora/StarTrekTNG/corpus.json", wordPerson)
		if err != nil {
			return fmt.Errorf("error loading corpus: %w", err)
		}

		rand.Seed(time.Now().UnixNano())
		chain := markov.NewChain(2)
		chain.BuildFromLines(cor.Lines)
		text := chain.Generate(wordLength)
		fmt.Println(text)
		return nil
	},
}

func init() {
	generateWordsCmd.Flags().StringVarP(&wordCorpus, "corpus", "c", "", "path to the corpus file")
	generateWordsCmd.Flags().StringVarP(&wordPerson, "person", "p", "", "person to base the generated text from")
	generateWordsCmd.Flags().IntVarP(&wordLength, "length", "l", 10, "number of words to generate")
	generateWordsCmd.MarkFlagRequired("corpus")
	rootCmd.AddCommand(generateWordsCmd)
}
