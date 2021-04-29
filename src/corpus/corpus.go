package corpus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Corpus is simply a collection of strings
type Corpus struct {
	Lines []string
}

// Load builds a corpus from a json file. This file attributes lines to specific
// "people". The corpus can be filtered to only lines by a specific "person". If
// person is an empty string, the corpus will not be filtered.
func Load(filename string, person string) (*Corpus, []string, error) {

	type line struct {
		Line   string
		Person string
	}

	lines := []*line{}
	people := map[string]bool{}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading corpus: %w", err)
	}

	err = json.Unmarshal(bytes, &lines)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading corpus: %w", err)
	}

	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("corpus contains no lines")
	}

	lineStrs := []string{}
	for _, line := range lines {
		people[strings.ToLower(line.Person)] = true
		if person == "" || strings.ToLower(person) == strings.ToLower(line.Person) {
			lineStrs = append(lineStrs, line.Line)
		}
	}

	if len(lineStrs) == 0 {
		return nil, nil, fmt.Errorf("person %s not found in corpus", person)
	}

	peopleStrs := []string{}
	for person := range people {
		peopleStrs = append(peopleStrs, person)
	}
	sort.Strings(peopleStrs)

	return &Corpus{Lines: lineStrs}, peopleStrs, nil
}
