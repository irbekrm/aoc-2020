package days

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	front, back, right, left rune = 70, 66, 82, 76
)

func Five() {
	b, err := ioutil.ReadFile("data/day5.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	bp := parsePasses(s)
	highest := highestIndex(bp)
	fmt.Printf("highest seat ID: %d\n", highest)
}

type boardingPass struct {
	row, column, id int
}

func parsePasses(s []string) []boardingPass {
	boardingPasses := make([]boardingPass, len(s))
	for key, val := range s {
		bp := boardingPass{}
		runes := []rune(val)
		firstIndex, lastIndex := 0, 127
		for i := 0; i < 7; i++ {
			if runes[i] == front {
				lastIndex = firstIndex + (lastIndex-firstIndex)/2
			}
			if runes[i] == back {
				firstIndex = firstIndex + (((lastIndex - firstIndex) / 2) + 1)
			}
		}
		bp.row = firstIndex
		firstIndex, lastIndex = 0, 7
		for i := 7; i < 10; i++ {
			if runes[i] == left {
				lastIndex = firstIndex + (lastIndex-firstIndex)/2
			}
			if runes[i] == right {
				firstIndex = firstIndex + (((lastIndex - firstIndex) / 2) + 1)
			}
		}
		bp.column = firstIndex
		bp.id = bp.row*8 + bp.column
		boardingPasses[key] = bp
	}
	return boardingPasses
}

func highestIndex(bp []boardingPass) int {
	highest := -1
	for _, val := range bp {
		if val.id > highest {
			highest = val.id
		}
	}
	return highest
}
