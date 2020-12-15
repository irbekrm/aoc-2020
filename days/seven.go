package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

const (
	bagsRegexp       = `^([a-z]+\s[a-z]+)|\d+\s([a-z]+\s[a-z]+)`
	bagsParentRegexp = `^[a-z]+\s[a-z]+`
	bagsChildRegexp  = `(\d+)\s([a-z]+\s[a-z]+)\sbag`
	shinyGoldBags    = "shiny gold"
)

var closure func([]string)

func Seven() {
	b, err := ioutil.ReadFile("data/day7.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	//bags := mapParents(s)
	//sum := countShinyGoldBags(bags)
	//fmt.Printf("sum: %v\n", sum)
	children := mapChildren(s)
	//fmt.Printf("%#v\n", children)
	sum := inShinyGold(children)
	fmt.Printf("sum: %v\n", sum)
}
func inShinyGold(allBags map[string][][]interface{}) int {
	var sum int
	ch := allBags[shinyGoldBags]
	for _, child := range ch {
		sum += countChildren(child[1].(int), child, allBags)
	}
	return sum
}

func countChildren(count int, children []interface{}, allBags map[string][][]interface{}) int {
	var sum int
	ch := allBags[children[0].(string)]
	for _, child := range ch {
		sum += countChildren(count*child[1].(int), child, allBags)
	}
	if len(ch) == 0 {
		return count
	}
	return sum + count
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

func mapChildren(s []string) map[string][][]interface{} {
	bags := make(map[string][][]interface{}, len(s))
	parentRe := regexp.MustCompile(bagsParentRegexp)
	childRe := regexp.MustCompile(bagsChildRegexp)
	for _, val := range s {
		parentMatch := parentRe.FindAllString(val, -1)
		childMatch := childRe.FindAllStringSubmatch(val, -1)
		parent := parentMatch[0]
		for _, child := range childMatch {
			count, err := strconv.Atoi(child[1])
			if err != nil {
				log.Fatalf("cannot convert %s to int: %v", child[1], err)
			}
			children := []interface{}{child[2], count}
			bags[parent] = append(bags[parent], children)
		}
	}
	return bags
}
