package rhymes

import (
	"fmt"
	"github.com/verkestk/goetry/src/corpus"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func Test_Load(t *testing.T) {
	cor, _, err := corpus.Load("../corpus/test_corpus.json", "")
	if err != nil {
		t.Errorf("Error loading corpus: %w", err)
	}
	if cor == nil {
		t.Errorf("Corpus is nil")
	}

	rhmr, err := Load("test_dictionary.txt", cor)
	if err != nil {
		t.Errorf("Error loading pronunciation dictionary: %w", err)
	}
	if rhmr == nil {
		t.Errorf("rhymer is nil")
	}
}

func Test_rhymer_Pronunciations(t *testing.T) {
	cor, _, _ := corpus.Load("../corpus/test_corpus.json", "")
	rhmr, _ := Load("test_dictionary.txt", cor)

	word := "beerbelly"
	pronunciations := rhmr.Pronunciations(word)
	if len(pronunciations) != 0 {
		t.Errorf("Expected 0 pronuncations for \"%s\", got %d", word, len(pronunciations))
	}

	word = "accidents"
	expectedPronunciations := [][]string{[]string{"AE1", "K", "S", "AH0", "D", "AH0", "N", "T", "S"}}
	pronunciations = rhmr.Pronunciations(word)
	if len(pronunciations) != len(expectedPronunciations) {
		t.Errorf("Expected %d pronuncations for \"%s\", got %d", len(expectedPronunciations), word, len(pronunciations))
	}
	if !reflect.DeepEqual(pronunciations, expectedPronunciations) {
		fmt.Println("expected:", expectedPronunciations)
		fmt.Println("actual:", pronunciations)
		t.Errorf("Unexpected pronuncations for \"%s\"", word)
	}

	word = "ACCIDENTS"
	expectedPronunciations = [][]string{[]string{"AE1", "K", "S", "AH0", "D", "AH0", "N", "T", "S"}}
	pronunciations = rhmr.Pronunciations(word)
	if len(pronunciations) != len(expectedPronunciations) {
		t.Errorf("Expected %d pronuncations for \"%s\", got %d", len(expectedPronunciations), word, len(pronunciations))
	}
	if !reflect.DeepEqual(pronunciations, expectedPronunciations) {
		fmt.Println("expected:", expectedPronunciations)
		fmt.Println("actual:", pronunciations)
		t.Errorf("Unexpected pronuncations for \"%s\"", word)
	}

	word = "am"
	expectedPronunciations = [][]string{[]string{"AM", "AE1", "M"}, []string{"EY1", "EH1", "M"}}
	pronunciations = rhmr.Pronunciations(word)
	if len(pronunciations) != len(expectedPronunciations) {
		t.Errorf("Expected %d pronuncations for \"%s\", got %d", len(expectedPronunciations), word, len(pronunciations))
	}

	word = "the"
	expectedPronunciations = [][]string{[]string{"DH", "AH0"}, []string{"DH", "AH1"}, []string{"DH", "IY0"}}
	pronunciations = rhmr.Pronunciations(word)
	if len(pronunciations) != len(expectedPronunciations) {
		t.Errorf("Expected %d pronuncations for \"%s\", got %d", len(expectedPronunciations), word, len(pronunciations))
	}

	word = "when"
	expectedPronunciations = [][]string{[]string{"W", "EH1", "N"}, []string{"HH", "W", "EH1", "N"}, []string{"W", "IH1", "N"}, []string{"HH", "W", "IH1", "N"}}
	pronunciations = rhmr.Pronunciations(word)
	if len(pronunciations) != len(expectedPronunciations) {
		t.Errorf("Expected %d pronuncations for \"%s\", got %d", len(expectedPronunciations), word, len(pronunciations))
	}
}

func Test_rhymer_Rhymes(t *testing.T) {
	wordsStrength1 := map[string][][]string{
		"a": [][]string{
			[]string{"hallelujah", "the", "to"},
			[]string{"away"},
		},
		"accidents": [][]string{
			[]string{"incidents"},
		},
		"ACCIDENTS": [][]string{
			[]string{"incidents"},
		},
		"al": [][]string{
			[]string{"pal"},
		},
		"allegations": [][]string{
			[]string{},
		},
		"alley": [][]string{
			[]string{"be", "betty", "family", "family", "infinity", "maybe", "opportunity", "poly", "the"},
		},
		"am": [][]string{
			[]string{},
			[]string{},
		},
		"amen": [][]string{
			[]string{"when", "when"},
			[]string{"when", "when"},
		},
		"anymore": [][]string{
			[]string{"door", "your"},
		},
		"are": [][]string{
			[]string{"far"},
			[]string{"architecture", "mr"},
		},
		"attention": [][]string{
			[]string{"can", "doesn't", "foreign", "foreign", "redemption", "redemption"},
		},
		"be": [][]string{
			[]string{"he", "me"},
			[]string{"alley", "betty", "family", "family", "infinity", "maybe", "opportunity", "poly", "the"},
		},
		"by": [][]string{
			[]string{"die", "i", "my", "why", "why"},
		},
		"can": [][]string{
			[]string{"man", "span"},
			[]string{"attention", "doesn't", "foreign", "foreign", "redemption", "redemption"},
		},
		"cattle": [][]string{
			[]string{"little", "middle", "model"},
		},
		"doesn't": [][]string{
			[]string{},
			[]string{"attention", "can", "foreign", "foreign", "redemption", "redemption"},
		},
		"family": [][]string{
			[]string{"alley", "be", "betty", "infinity", "maybe", "opportunity", "poly", "the"},
			[]string{"alley", "be", "betty", "infinity", "maybe", "opportunity", "poly", "the"},
		},
		"foreign": [][]string{
			[]string{"attention", "can", "doesn't", "redemption", "redemption"},
			[]string{"attention", "can", "doesn't", "redemption", "redemption"},
		},
		"get": [][]string{
			[]string{},
			[]string{"lit"},
		},
		"in": [][]string{
			[]string{},
			[]string{"when", "when"},
		},
		"the": [][]string{
			[]string{"a", "hallelujah", "to"},
			[]string{},
			[]string{"alley", "be", "betty", "family", "family", "infinity", "maybe", "opportunity", "poly"},
		},
		"to": [][]string{
			[]string{"you"},
			[]string{},
			[]string{"a", "hallelujah", "the"},
		},
		"when": [][]string{
			[]string{"amen", "amen"},
			[]string{"amen", "amen"},
			[]string{"in"},
			[]string{"in"},
		},
		"with": [][]string{
			[]string{},
			[]string{},
			[]string{},
			[]string{},
		},
		"whoa": [][]string{
			[]string{"know", "no", "photo", "so"},
			[]string{"know", "no", "photo", "so"},
			[]string{"know", "no", "photo", "so"},
		},
		"why": [][]string{
			[]string{"by", "die", "i", "my"},
			[]string{"by", "die", "i", "my"},
		},
	}

	wordsStrength2 := map[string][][]string{
		"a": [][]string{
			[]string{},
			[]string{},
		},
		"accidents": [][]string{
			[]string{"incidents"},
		},
		"ACCIDENTS": [][]string{
			[]string{"incidents"},
		},
	}

	cor, _, _ := corpus.Load("../corpus/test_corpus.json", "")
	rhmr, _ := Load("test_dictionary.txt", cor)

	for word, expectedRhymeGroups := range wordsStrength1 {
		pronunciations := rhmr.Pronunciations(word)
		if len(expectedRhymeGroups) != len(pronunciations) {
			fmt.Println("expected:", expectedRhymeGroups)
			fmt.Println("actual:", pronunciations)
			t.Fatalf("Expected %d pronunciations for \"%s\", got %d", len(expectedRhymeGroups), word, len(pronunciations))
		}
		for indexP, pronunciation := range pronunciations {
			rhymes := rhmr.Rhymes(word, pronunciation, 1)
			if len(expectedRhymeGroups[indexP]) != len(rhymes) {
				actual := []string{}
				for _, rhyme := range rhymes {
					actual = append(actual, rhyme.Word)
				}
				fmt.Println("expected:", expectedRhymeGroups[indexP])
				fmt.Println("actual:", actual)
				t.Fatalf("Expected %d rhymes for \"%s\" (%s), got %d", len(expectedRhymeGroups[indexP]), word, strings.Join(pronunciation, " "), len(rhymes))
			}

			for indexR, rhyme := range rhymes {
				if expectedRhymeGroups[indexP][indexR] != rhyme.Word {
					t.Errorf("Expected rhyme index [%d][%d] for \"%s\" to be \"%s\", got \"%s\"", indexP, indexR, word, expectedRhymeGroups[indexP][indexR], rhyme.Word)
				}
				if rhyme.Strength < 1 {
					t.Errorf("Expected rhyme index [%d][%d] for \"%s\" to have strength at least 1, got %d", indexP, indexR, word, rhyme.Strength)
				}
			}
		}
	}

	for word, expectedRhymeGroups := range wordsStrength2 {
		pronunciations := rhmr.Pronunciations(word)
		if len(expectedRhymeGroups) != len(pronunciations) {
			fmt.Println("expected:", expectedRhymeGroups)
			fmt.Println("actual:", pronunciations)
			t.Fatalf("Expected %d pronunciations for \"%s\", got %d", len(expectedRhymeGroups), word, len(pronunciations))
		}
		for indexP, pronunciation := range pronunciations {
			rhymes := rhmr.Rhymes(word, pronunciation, 2)
			if len(expectedRhymeGroups[indexP]) != len(rhymes) {
				actual := []string{}
				for _, rhyme := range rhymes {
					actual = append(actual, rhyme.Word)
				}
				fmt.Println("expected:", expectedRhymeGroups[indexP])
				fmt.Println("actual:", actual)
				t.Fatalf("Expected %d rhymes for \"%s\" (%s), got %d", len(expectedRhymeGroups[indexP]), word, strings.Join(pronunciation, " "), len(rhymes))
			}

			for indexR, rhyme := range rhymes {
				if expectedRhymeGroups[indexP][indexR] != rhyme.Word {
					t.Errorf("Expected rhyme index [%d][%d] for \"%s\" to be \"%s\", got \"%s\"", indexP, indexR, word, expectedRhymeGroups[indexP][indexR], rhyme.Word)
				}
				if rhyme.Strength < 2 {
					t.Errorf("Expected rhyme index [%d][%d] for \"%s\" to have strength at least 2, got %d", indexP, indexR, word, rhyme.Strength)
				}
			}
		}
	}
}

func Test_rhymer_UnknownPronunciations(t *testing.T) {
	cor, _, _ := corpus.Load("../corpus/test_corpus.json", "")
	rhmr, _ := Load("test_dictionary.txt", cor)

	unknown := rhmr.UnknownPronunciations()
	if len(unknown) != 5 {
		t.Errorf("Expected 5 unknown words, got %d", len(unknown))
	}

	index := 0
	word := "beerbelly"
	if unknown[index] != word {
		t.Errorf("Expected \"%s\" at index %d, got \"%s\"", word, index, unknown[index])
	}

	index = 1
	word = "bonedigger"
	if unknown[index] != word {
		t.Errorf("Expected \"%s\" at index %d, got \"%s\"", word, index, unknown[index])
	}

	index = 2
	word = "currency"
	if unknown[index] != word {
		t.Errorf("Expected \"%s\" at index %d, got \"%s\"", word, index, unknown[index])
	}

	index = 3
	word = "roly"
	if unknown[index] != word {
		t.Errorf("Expected \"%s\" at index %d, got \"%s\"", word, index, unknown[index])
	}

	index = 4
	word = "scatterings"
	if unknown[index] != word {
		t.Errorf("Expected \"%s\" at index %d, got \"%s\"", word, index, unknown[index])
	}
}

// less important ()
func Test_getPronunciationFromDictionary(t *testing.T) {
	line := "WORD  PHONEME1 PHONEME2"
	expectedWord := "word"
	expectedPhonemes := []string{"PHONEME1", "PHONEME2"}
	word, phonemes := getPronunciationFromDictionary(line)
	if word != expectedWord {
		t.Errorf("expected word \"%s\", got \"%s\"", expectedWord, word)
	}
	if !reflect.DeepEqual(expectedPhonemes, phonemes) {
		fmt.Println("expected", expectedPhonemes)
		fmt.Println("actual", phonemes)
		t.Errorf("unexpected phonemes")
	}

	line = "WORD(1)  PHONEME1 PHONEME2"
	expectedWord = "word"
	expectedPhonemes = []string{"PHONEME1", "PHONEME2"}
	word, phonemes = getPronunciationFromDictionary(line)
	if word != expectedWord {
		t.Errorf("expected word \"%s\", got \"%s\"", expectedWord, word)
	}
	if !reflect.DeepEqual(expectedPhonemes, phonemes) {
		fmt.Println("expected", expectedPhonemes)
		fmt.Println("actual", phonemes)
		t.Errorf("unexpected phonemes")
	}

	line = "INVALID"
	word, phonemes = getPronunciationFromDictionary(line)
	if word != "" {
		t.Errorf("expected empty word, got \"%s\"", word)
	}
	if len(phonemes) > 0 {
		fmt.Println("expected", []string{})
		fmt.Println("actual", phonemes)
		t.Errorf("unexpected phonemes")
	}
}

func Test_rhymeStrength(t *testing.T) {

	// strength 4
	word1 := "astrophotography"
	word2 := "chromatography"
	pronunciation1 := []string{"AE2", "S", "T", "R", "OW0", "F", "AH0", "T", "AA1", "G", "R", "AH0", "F", "IY0"}
	pronunciation2 := []string{"K", "R", "OW0", "M", "AH0", "T", "AA1", "G", "R", "AH0", "F", "IY0"}
	expectedStrength := 4
	strength := rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// strength 3
	word1 = "chronicle"
	word2 = "monical"
	pronunciation1 = []string{"K", "R", "AA1", "N", "IH0", "K", "AH0", "L"}
	pronunciation2 = []string{"M", "AA1", "N", "IH0", "K", "AH0", "L"}
	expectedStrength = 3
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// strength 2
	word1 = "accident"
	word2 = "incident"
	pronunciation1 = []string{"AE1", "K", "S", "AH0", "D", "AH0", "N", "T"}
	pronunciation2 = []string{"IH1", "N", "S", "AH0", "D", "AH0", "N", "T"}
	expectedStrength = 2
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// strength 1
	word1 = "accessorize"
	word2 = "acidifies"
	pronunciation1 = []string{"AE0", "K", "S", "EH1", "S", "ER0", "AY2", "Z"}
	pronunciation2 = []string{"AH0", "S", "IH1", "D", "AH0", "F", "AY2", "Z"}
	expectedStrength = 1
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// strength 1
	word1 = "beaudreau"
	word2 = "bordeaux"
	pronunciation1 = []string{"B", "OW2", "D", "R", "OW1"}
	pronunciation2 = []string{"B", "AO0", "R", "D", "OW1"}
	expectedStrength = 1
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// strength 0
	word1 = "caffeine"
	word2 = "calculus"
	pronunciation1 = []string{"K", "AE0", "F", "IY1", "N"}
	pronunciation2 = []string{"K", "AE1", "L", "K", "Y", "AH0", "L", "AH0", "S"}
	expectedStrength = 0
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// different words, same phonemes
	word1 = "read"
	word2 = "reed"
	pronunciation1 = []string{"R", "IY1", "D"}
	pronunciation2 = []string{"R", "IY1", "D"}
	expectedStrength = -1
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// same word, different phonemes
	word1 = "lead"
	word2 = "lead"
	pronunciation1 = []string{"L", "EH1", "D"}
	pronunciation2 = []string{"L", "IY1", "D"}
	expectedStrength = -1
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// emphasis should match
	word1 = "emph1"
	word2 = "emph2"
	pronunciation1 = []string{"Y", "A1"}
	pronunciation2 = []string{"Z", "A2"}
	expectedStrength = 1
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}

	// emphasis should not match
	word1 = "emph0"
	word2 = "emph1"
	pronunciation1 = []string{"X", "A0"}
	pronunciation2 = []string{"Y", "A1"}
	expectedStrength = 0
	strength = rhymeStrength(word1, word2, pronunciation1, pronunciation2)
	if expectedStrength != strength {
		t.Errorf("Expected rhyme strength %d for \"%s\" and \"%s\", got %d", expectedStrength, word1, word2, strength)
	}
}

func Test_getRhymeSyllables(t *testing.T) {
	phonemes := []string{"A0"}
	expectedSyllables := []string{"A0"}
	syllables := getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"A0", "B"}
	expectedSyllables = []string{"A0B"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"B", "A0"}
	expectedSyllables = []string{"A0"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"B", "A0", "B"}
	expectedSyllables = []string{"A0B"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"A0", "B", "A0"}
	expectedSyllables = []string{"A0B", "A0"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"A0", "B", "A0", "B"}
	expectedSyllables = []string{"A0B", "A0B"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"B", "A0", "B", "A0"}
	expectedSyllables = []string{"A0B", "A0"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"B", "A0", "B", "A0", "B"}
	expectedSyllables = []string{"A0B", "A0B"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"A0", "A0"}
	expectedSyllables = []string{"A0", "A0"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"A0", "A0", "B", "B"}
	expectedSyllables = []string{"A0", "A0BB"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"B", "B", "A0", "A0", "B", "B", "A0", "A0"}
	expectedSyllables = []string{"A0", "A0BB", "A0", "A0"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"B", "B", "A0", "A0", "B", "B", "A0", "A0", "B", "B"}
	expectedSyllables = []string{"A0", "A0BB", "A0", "A0BB"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"A0", "A1", "A2"}
	expectedSyllables = []string{"A0", "A1", "A1"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}

	phonemes = []string{"S", "UW2", "P", "ER0", "K", "AE2", "L", "AH0", "F", "R", "AE1", "JH", "AH0", "L", "IH2", "S", "T", "IH0", "K", "EH2", "K", "S", "P", "IY0", "AE2", "L", "AH0", "D", "OW1", "SH", "AH0", "S"}
	expectedSyllables = []string{"UW1P", "ER0K", "AE1L", "AH0FR", "AE1JH", "AH0L", "IH1ST", "IH0K", "EH1KSP", "IY0", "AE1L", "AH0D", "OW1SH", "AH0S"}
	syllables = getRhymeSyllables(phonemes)
	if !reflect.DeepEqual(expectedSyllables, syllables) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedSyllables)
		fmt.Println("actual:", syllables)
		t.Errorf("unexpected syllables")
	}
}

func Test_indexFirstVowel(t *testing.T) {
	phonemes := []string{"A0", "B", "C", "D"}
	expectedIndex := 0
	index := indexFirstVowel(phonemes)
	if expectedIndex != index {
		fmt.Println("phonemes:", phonemes)
		t.Errorf("expected %d first vowel index, got %d", expectedIndex, index)
	}

	phonemes = []string{"A", "B1", "C", "D"}
	expectedIndex = 1
	index = indexFirstVowel(phonemes)
	if expectedIndex != index {
		fmt.Println("phonemes:", phonemes)
		t.Errorf("expected %d first vowel index, got %d", expectedIndex, index)
	}

	phonemes = []string{"A", "B", "C2", "D"}
	expectedIndex = 2
	index = indexFirstVowel(phonemes)
	if expectedIndex != index {
		fmt.Println("phonemes:", phonemes)
		t.Errorf("expected %d first vowel index, got %d", expectedIndex, index)
	}

	phonemes = []string{"A", "B", "C", "D3"}
	expectedIndex = 3
	index = indexFirstVowel(phonemes)
	if expectedIndex != index {
		fmt.Println("phonemes:", phonemes)
		t.Errorf("expected %d first vowel index, got %d", expectedIndex, index)
	}

	phonemes = []string{"A0", "B1", "C2", "D3"}
	expectedIndex = 0
	index = indexFirstVowel(phonemes)
	if expectedIndex != index {
		fmt.Println("phonemes:", phonemes)
		t.Errorf("expected %d first vowel index, got %d", expectedIndex, index)
	}
}

func Test_isVowelPhoneme(t *testing.T) {
	phoneme := "A"
	expected := false
	actual := isVowelPhoneme(phoneme)
	if expected != actual {
		t.Errorf("Expected isVowelPhoneme %v for \"%s\", got %v", expected, phoneme, actual)
	}

	phoneme = "A1"
	expected = true
	actual = isVowelPhoneme(phoneme)
	if expected != actual {
		t.Errorf("Expected isVowelPhoneme %v for \"%s\", got %v", expected, phoneme, actual)
	}

	phoneme = "1A"
	expected = false
	actual = isVowelPhoneme(phoneme)
	if expected != actual {
		t.Errorf("Expected isVowelPhoneme %v for \"%s\", got %v", expected, phoneme, actual)
	}
}

func Test_normalizeEmphasis(t *testing.T) {
	phonemes := []string{"S", "UW2", "P", "ER0", "K", "AE2", "L", "AH0", "F", "R", "AE1", "JH", "AH0", "L", "IH2", "S", "T", "IH0", "K", "EH2", "K", "S", "P", "IY0", "AE2", "L", "AH0", "D", "OW1", "SH", "AH0", "S"}
	expectedNormalized := []string{"S", "UW1", "P", "ER0", "K", "AE1", "L", "AH0", "F", "R", "AE1", "JH", "AH0", "L", "IH1", "S", "T", "IH0", "K", "EH1", "K", "S", "P", "IY0", "AE1", "L", "AH0", "D", "OW1", "SH", "AH0", "S"}
	normalized := normalizeEmphasis(phonemes)
	if !reflect.DeepEqual(expectedNormalized, normalized) {
		fmt.Println("phonemes:", phonemes)
		fmt.Println("expected:", expectedNormalized)
		fmt.Println("actual:", normalized)
		t.Errorf("unexpected emphasis")
	}

	syllables := []string{"UW2P", "ER0K", "AE2L", "AH0FR", "AE1JH", "AH0L", "IH2ST", "IH0K", "EH2KSP", "IY0", "AE2L", "AH0D", "OW1SH", "AH0S"}
	expectedNormalized = []string{"UW1P", "ER0K", "AE1L", "AH0FR", "AE1JH", "AH0L", "IH1ST", "IH0K", "EH1KSP", "IY0", "AE1L", "AH0D", "OW1SH", "AH0S"}
	normalized = normalizeEmphasis(syllables)
	if !reflect.DeepEqual(expectedNormalized, normalized) {
		fmt.Println("syllables:", syllables)
		fmt.Println("expected:", expectedNormalized)
		fmt.Println("actual:", normalized)
		t.Errorf("unexpected emphasis")
	}
}

func Test_byStrengthDesc(t *testing.T) {

	rhyme1 := &Rhyme{Word: "what", Strength: 1}
	rhyme2 := &Rhyme{Word: "noise", Strength: 2}
	rhyme3 := &Rhyme{Word: "annoys", Strength: 3}
	rhyme4 := &Rhyme{Word: "a", Strength: 1}
	rhyme5 := &Rhyme{Word: "noisy", Strength: 2}
	rhyme6 := &Rhyme{Word: "oyster", Strength: 3}

	rhymes := []*Rhyme{rhyme1, rhyme2, rhyme3, rhyme4, rhyme5, rhyme6}

	sort.Sort(byStrengthDesc(rhymes))
	if rhymes[0].Word != "annoys" || rhymes[1].Word != "oyster" || rhymes[2].Word != "noise" || rhymes[3].Word != "noisy" || rhymes[4].Word != "a" || rhymes[5].Word != "what" {
		expectedOrder := []string{"annoys", "oyster", "noise", "noisy", "a", "what"}
		actualOrder := []string{}
		for _, rhyme := range rhymes {
			actualOrder = append(actualOrder, rhyme.Word)
		}
		fmt.Println("excpted:", expectedOrder)
		fmt.Println("actual:", actualOrder)
		t.Errorf("unexpected ordering")
	}
}
