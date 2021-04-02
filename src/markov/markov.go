// Adapted from https://golang.org/doc/codewalk/markov/

/*
Generating random text: a Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix. Consider this text:

	I am not a number! I am a free man!

Our Markov chain algorithm would arrange this text into this set of prefixes
and suffixes, or "chain": (This table assumes a prefix length of two words.)

	Prefix       Suffix

	"" ""        I
	"" I         am
	I am         a
	I am         not
	a free       man!
	am a         free
	am not       a
	a number!    I
	number! I    am
	not a        number!

To generate text using this table we select an initial prefix ("I am", for
example), choose one of the suffixes associated with that prefix at random
with probability determined by the input statistics ("a"),
and then create a new prefix by removing the first word from the prefix
and appending the suffix (making the new prefix is "am a"). Repeat this process
until we can't find any suffixes for the current prefix or we exceed the word
limit. (The word limit is necessary as the chain table may contain cycles.)

Our version of this program reads text from standard input, parsing it into a
Markov chain, and writes generated text to standard output.
The prefix and output lengths can be specified using the -prefix and -words
flags on the command-line.
*/
package markov

import (
	"math/rand"
	"strings"
)

// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// Build reads text and parses it into prefixes and suffixes that are stored in
// Chain.
func (c *Chain) Build(tokens []string) {
	p := make(Prefix, c.prefixLen)
	for _, token := range tokens {
		key := p.String()
		c.chain[key] = append(c.chain[key], token)
		p.Shift(token)
	}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) BuildFromLines(lines []string) {
	for _, line := range lines {
		p := make(Prefix, c.prefixLen)
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			key := p.String()
			c.chain[key] = append(c.chain[key], token)
			p.Shift(token)
		}
	}
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func (c *Chain) GenerateSentences(length int) string {
	sentences := []string{}

	for len(sentences) < length {
		s := c.generateSentences(length - len(sentences))
		sentences = append(sentences, s...)
	}

	return strings.Join(sentences, " ")
}

// attemps to generate n sentences, but might not get all the way there
func (c *Chain) generateSentences(length int) []string {
	p := make(Prefix, c.prefixLen)
	var words []string
	var sentences []string

	for {
		choices := c.chain[p.String()]

		// if you run out of chain, stop
		if len(choices) == 0 {
			break
		}

		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)

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
