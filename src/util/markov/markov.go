package markov

import (
	"strings"

	"github.com/verkestk/markovokram"
)

// GenerateWords returns a string of at most _length_ words generated from
// Chain.
func GenerateWords(chain *markovokram.Chain, length int) string {
	tokens := []string{}
	generation := chain.GenerateForward()

	next := generation.Next()
	for len(tokens)+1 < length && next != "" {
		tokens = append(tokens, next)
		next = generation.Next()
	}

	tokens = append(tokens, next)

	return strings.Join(tokens, " ")
}

// GenerateSentences generates _length_ sentences - defining a sentences by a
// generated sequence ending in a ".". It's possible this function will never
// complete if there are no tokens ending in "." in the corpus.
func GenerateSentences(chain *markovokram.Chain, length int) string {
	sentences := []string{}

	for len(sentences) < length {
		s := generateSentences(chain, length-len(sentences))
		sentences = append(sentences, s...)
	}

	return strings.Join(sentences, " ")
}

// attemps to generate n sentences, but might not get all the way there
func generateSentences(chain *markovokram.Chain, length int) []string {
	var words []string
	var sentences []string

	generation := chain.GenerateForward()

	for {
		next := generation.Next()

		// if you run out of chain, stop
		if next == "" {
			break
		}

		words = append(words, next)

		if next[len(next)-1:] == "." || next[len(next)-1:] == "?" || next[len(next)-1:] == "!" {
			sentences = append(sentences, strings.Join(words, " "))
			words = []string{}

			if len(sentences) == length {
				break
			}
		}
	}
	return sentences
}
