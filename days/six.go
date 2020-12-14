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
	fmt.Printf("sum is %d\n", countAnswers(forms))
}

func countAnswers(forms []map[rune]int) int {
	var sum int
	for _, f := range forms {
		sum += len(f)
	}
	return sum
}

func parseForms(s []string) []map[rune]int {
	forms := make([]map[rune]int, len(s))
	for key, val := range s {
		m := make(map[rune]int)
		chars := strings.ReplaceAll(val, "\n", "")
		for _, ch := range chars {
			m[ch]++
		}
		forms[key] = m
	}
	return forms
}
