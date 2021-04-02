package corpus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Corpus struct {
	Lines []string
}

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

	lineStrs := []string{}
	for _, line := range lines {
		people[strings.ToLower(line.Person)] = true
		if person == "" || strings.ToLower(person) == strings.ToLower(line.Person) {
			lineStrs = append(lineStrs, line.Line)
		}
	}

	peopleStrs := []string{}
	for person, _ := range people {
		peopleStrs = append(peopleStrs, person)
	}
	sort.Strings(peopleStrs)

	return &Corpus{Lines: lineStrs}, peopleStrs, nil
}
