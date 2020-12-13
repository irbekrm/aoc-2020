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
	fmt.Printf("my seat id: %d\n", findEmptySeat(bp))
}

func findEmptySeat(boardingPasses map[int]*boardingPass) int {
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			bp := boardingPass{j, i}
			id := bp.calcId()
			if boardingPasses[id] == nil {
				if boardingPasses[id+1] != nil && boardingPasses[id-1] != nil {
					return id
				}
			}
		}
	}
	return -1
}

type boardingPass struct {
	column, row int
}

func (bp *boardingPass) calcId() int {
	return bp.row*8 + bp.column
}

func parsePasses(s []string) map[int]*boardingPass {
	boardingPasses := make(map[int]*boardingPass, len(s))
	for _, val := range s {
		bp := &boardingPass{}
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
		id := bp.calcId()
		boardingPasses[id] = bp
	}
	return boardingPasses
}

func highestIndex(bp map[int]*boardingPass) int {
	highest := -1
	for key := range bp {
		if key > highest {
			highest = key
		}
	}
	return highest
}
