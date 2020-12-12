package days

import (
	"fmt"
	"io/ioutil"
	"log"
)

const tree rune = 35

func Three() {
	b, err := ioutil.ReadFile("data/day3.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	mtrx := newMatrix(s)
	// traverse 1 1
	x := mtrx.traverse(1, 1)
	// traverse 3 1
	y := mtrx.traverse(3, 1)
	// traverse 5 1
	z := mtrx.traverse(5, 1)
	// traverse 7 1
	k := mtrx.traverse(7, 1)
	// taverse 1 2
	t := mtrx.traverse(1, 2)
	fmt.Println(x * y * z * k * t)
}

type matrix struct {
	m [][]bool
}

// traverse travels along matrix, returns the number of trees encountered
func (mtrx *matrix) traverse(rightStep, downStep int) int {
	trees := 0
	for right, down := 0, 0; down < len(mtrx.m); {
		index := right % len(mtrx.m[0])
		if mtrx.m[down][index] {
			trees++
		}
		right += rightStep
		down += downStep
	}
	return trees
}

func newMatrix(input []string) *matrix {
	m := make([][]bool, len(input))
	for k := range m {
		m[k] = make([]bool, len(input[0]))
		for i, char := range input[k] {
			m[k][i] = (char == tree)
		}
	}
	return &matrix{m}
}
