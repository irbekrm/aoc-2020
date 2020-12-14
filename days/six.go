package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Six() {
	b, err := ioutil.ReadFile("data/day6.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n\n")
	forms := parseForms(s)
	fmt.Printf("sum is %d\n", countEveryonesAnswers(forms))
}

func countAnswers(forms []*form) int {
	var sum int
	for _, f := range forms {
		sum += len(f.forms)
	}
	return sum
}

func countEveryonesAnswers(forms []*form) int {
	var sum int
	for _, f := range forms {
		for _, val := range f.forms {
			if val == f.groupSize {
				sum++
			}
		}
	}
	return sum
}

type form struct {
	forms     map[rune]int
	groupSize int
}

func parseForms(s []string) []*form {
	forms := make([]*form, len(s))
	for key, val := range s {
		m := make(map[rune]int)
		numOfGroups := strings.Count(val, "\n") + 1
		chars := strings.ReplaceAll(val, "\n", "")
		for _, ch := range chars {
			m[ch]++
		}
		forms[key] = &form{m, numOfGroups}
	}
	return forms
}
