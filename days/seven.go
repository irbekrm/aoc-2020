package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

const (
	bagsRegexp    = `^([a-z]+\s[a-z]+)|\d+\s([a-z]+\s[a-z]+)`
	shinyGoldBags = "shiny gold"
)

var closure func([]string)

func Seven() {
	b, err := ioutil.ReadFile("data/day7.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	bags := mapParents(s)
	sum := countShinyGoldBags(bags)
	fmt.Printf("sum: %v\n", sum)
}

func countShinyGoldBags(bags map[string][]string) int {
	predecessors := make(map[string]bool)
	closure = func(s []string) {
		for _, bag := range s {
			if !predecessors[bag] {
				predecessors[bag] = true
				closure(bags[bag])
			}
		}
	}
	closure(bags[shinyGoldBags])
	return len(predecessors)
}

func mapParents(s []string) map[string][]string {
	bags := make(map[string][]string, len(s))
	re := regexp.MustCompile(bagsRegexp)
	for _, val := range s {
		match := re.FindAllStringSubmatch(val, -1)
		parent := match[0][1]
		for i := 1; i < len(match); i++ {
			bags[match[i][2]] = append(bags[match[i][2]], parent)
		}
	}
	return bags
}
