package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Eight() {
	b, err := ioutil.ReadFile("data/day8.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	instructions := parseInstructions(s)
	accumulator := stepThroughOnce(instructions)
	fmt.Printf("accumulator: %d\n", accumulator)
}

type instruction struct {
	op   string
	arg  int
	seen bool
}

func stepThroughOnce(instrs []*instruction) int {
	accumulator, index := 0, 0
	for {
		if instrs[index].seen {
			return accumulator
		}
		instrs[index].seen = true
		if instrs[index].op == "jmp" {
			index += instrs[index].arg
			continue
		}
		if instrs[index].op == "acc" {
			accumulator += instrs[index].arg
		}
		index++
	}
}

func parseInstructions(s []string) []*instruction {
	instructions := make([]*instruction, len(s))
	for key, value := range s {
		sl := strings.Split(value, " ")
		if len(sl) != 2 {
			log.Fatalf("%s could not be split on whitespace", value)
		}
		arg, err := strconv.Atoi(sl[1])
		if err != nil {
			log.Fatalf("error converting %s to int: %v", sl[1], err)
		}
		instructions[key] = &instruction{arg: arg, op: sl[0]}
	}
	return instructions
}
