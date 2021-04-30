package rhymes

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/verkestk/goetry/src/corpus"
)

// Rhyme is a word plus it's pronunciation
type Rhyme struct {
	// the word itself, as originally appearing in the corpus
	Word string

	// the pronunciation of the Word
	Pronunciation []string

	// the strength of the rhyme (roughly number of rhyming syllables)
	Strength int
}

// Rhymer provides functions for getting pronunciations, finding rhyming words
// from a corpus, and finding words in corpus missing pronunciation data.
type Rhymer struct {
	rhymes  map[string][]*Rhyme
	missing map[string]bool
}

type byStrengthDesc []*Rhyme

func (s byStrengthDesc) Len() int {
	return len(s)
}
func (s byStrengthDesc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byStrengthDesc) Less(i, j int) bool {
	if s[i].Strength == s[j].Strength {
		return s[i].Word < s[j].Word
	}
	return s[i].Strength > s[j].Strength
}

// Load creates a Rhymer based on a corpus, using a specific rhyming dictionary
// It finds rhymes
func Load(pronunciationDictionaryFilepath string, corpus *corpus.Corpus) (*Rhymer, error) {
	bytes, err := ioutil.ReadFile(pronunciationDictionaryFilepath)
	if err != nil {
		return nil, fmt.Errorf("error loading pronunciation dictionary: %w", err)
	}

	// map the pronunciations for everythingg in the dictionary for quick
	// reference
	pronunciationMap := map[string][][]string{}

	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		word, pronunciation := getPronunciationFromDictionary(line)
		_, ok := pronunciationMap[word]
		if !ok {
			pronunciationMap[word] = [][]string{pronunciation}
		} else {
			pronunciationMap[word] = append(pronunciationMap[word], pronunciation)
		}
	}

	// get all of the words from the corpus and save all of their pronunciations
	// in a *rhymer
	rhmr := &Rhymer{rhymes: make(map[string][]*Rhyme), missing: make(map[string]bool)}
	for _, line := range corpus.Lines {
		// tokenize by all non letter/numbers (excluding apostrophes)
		words := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
		})
		for _, word := range words {
			_, ok := rhmr.rhymes[strings.ToLower(word)]
			if !ok {
				rhymes := []*Rhyme{}
				prounciations, ok := pronunciationMap[strings.ToLower(word)]
				if ok {
					for _, pronunciation := range prounciations {
						rhymes = append(rhymes, &Rhyme{Word: strings.ToLower(word), Pronunciation: pronunciation})
					}
				} else {
					rhmr.missing[strings.ToLower(word)] = true
				}
				rhmr.rhymes[strings.ToLower(word)] = rhymes
			}
		}
	}

	return rhmr, nil
}

// Pronunciations provides the pronunciation of a word. Returns empty string for
// unknown words. A single word can have multiple pronunciations. Each
// pronunciation is represented by a string slice of phonemes.
func (r *Rhymer) Pronunciations(word string) [][]string {
	rhymes, ok := r.rhymes[strings.ToLower(word)]
	if ok {
		pronunciations := [][]string{}
		for _, rhyme := range rhymes {
			pronunciations = append(pronunciations, rhyme.Pronunciation)
		}

		return pronunciations
	}

	return nil
}

// Rhymes returns a list of Rhymes that match the word, ordered by strength of
// the rhyme (number of rhyming syllables).
func (r *Rhymer) Rhymes(word string, pronunciation []string, minStrength int) []*Rhyme {
	// TODO - efficiant algorithm for finding rhymes

	actualRhymes := []*Rhyme{}

	for _, rhymeList := range r.rhymes {
		for _, rhyme := range rhymeList {
			strength := rhymeStrength(word, rhyme.Word, pronunciation, rhyme.Pronunciation)
			if strength >= minStrength {
				actualRhymes = append(actualRhymes, &Rhyme{Word: rhyme.Word, Pronunciation: rhyme.Pronunciation, Strength: strength})
			}
		}
	}

	sort.Sort(byStrengthDesc(actualRhymes))

	return actualRhymes
}

// UnknownPronunciations returns all the words from the corpus that have no
// known pronunciation.
func (r *Rhymer) UnknownPronunciations() []string {
	missing := []string{}
	for m := range r.missing {
		missing = append(missing, m)
	}

	sort.Strings(missing)
	return missing
}

func getPronunciationFromDictionary(line string) (string, []string) {
	pieces := strings.Split(line, " ")

	if len(pieces) < 3 {
		return "", nil
	}

	word := strings.ToLower(pieces[0])
	pronunciation := pieces[2:]

	leftParenIndex := strings.Index(word, "(")
	if leftParenIndex > 0 {
		word = word[:leftParenIndex]
	}

	return word, pronunciation
}

func rhymeStrength(word1, word2 string, pronunciation1, pronunciation2 []string) int {
	if strings.ToLower(word1) == strings.ToLower(word2) || reflect.DeepEqual(normalizeEmphasis(pronunciation1), normalizeEmphasis(pronunciation2)) {
		return -1
	}

	rhymeSyllables1 := getRhymeSyllables(pronunciation1)
	rhymeSyllables2 := getRhymeSyllables(pronunciation2)

	if len(rhymeSyllables1) > len(rhymeSyllables2) {
		rhymeSyllables1 = rhymeSyllables1[len(rhymeSyllables1)-len(rhymeSyllables2):]
	} else if len(rhymeSyllables2) > len(rhymeSyllables1) {
		rhymeSyllables2 = rhymeSyllables2[len(rhymeSyllables2)-len(rhymeSyllables1):]
	}

	strength := 0
	for i := range rhymeSyllables1 {
		if rhymeSyllables1[len(rhymeSyllables1)-i-1] == rhymeSyllables2[len(rhymeSyllables2)-i-1] {
			strength++
		} else {
			break
		}
	}

	return strength
}

func getRhymeSyllables(pronunciation []string) []string {
	// list := C C V V C C V C => V VCC VC

	firstVowel := indexFirstVowel(pronunciation)
	if firstVowel == -1 {
		return nil
	}

	// chunk them by phoneme type (consonants/vowels)
	syllables := []string{}
	currentSyllable := ""
	for i := firstVowel; i < len(pronunciation); i++ {
		if (len(syllables) == 0 && len(currentSyllable) == 0) || !isVowelPhoneme(pronunciation[i]) {
			// if this is the first vowel, or any consonant, append to current syllable
			currentSyllable += pronunciation[i]
		} else {
			syllables = append(syllables, currentSyllable)
			currentSyllable = pronunciation[i]
		}
	}
	syllables = append(syllables, currentSyllable)

	return normalizeEmphasis(syllables)
}

func indexFirstVowel(pronunciation []string) int {
	for i, phoneme := range pronunciation {
		if isVowelPhoneme(phoneme) {
			return i
		}
	}

	return -1
}

func isVowelPhoneme(phoneme string) bool {
	lastRune, _ := utf8.DecodeLastRuneInString(phoneme)
	return unicode.IsNumber(lastRune)
}

func normalizeEmphasis(syllablesOrPhonemes []string) []string {
	normalized := []string{}
	for _, str := range syllablesOrPhonemes {
		normalized = append(normalized, strings.ReplaceAll(str, "2", "1"))
	}

	return normalized
}
