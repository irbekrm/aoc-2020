package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	empty    string = "L"
	occupied string = "#"
	floor    string = "."
)

func Eleven() {
	b, err := ioutil.ReadFile("data/day11.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	matrix := buildMatrix(s)
	changedMatrix, isChanged := changeSeating(matrix)
	for isChanged {
		changedMatrix, isChanged = changeSeating(changedMatrix)
	}
	occupiedSeats := countOccupiedSeats(changedMatrix)
	fmt.Printf("number of occupied seats is %d\n", occupiedSeats)
}

func buildMatrix(s []string) [][]string {
	matrix := make([][]string, len(s))
	for key, val := range s {
		row := strings.Split(val, "")
		matrix[key] = row
	}
	return matrix
}

func countOccupiedSeats(matrix [][]string) int {
	var occupiedSeats int
	for _, column := range matrix {
		for _, seat := range column {
			if seat == occupied {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func changeSeating(matrix [][]string) ([][]string, bool) {
	var changed bool
	changedSeats := make([][]string, len(matrix))
	for rowIndex, column := range matrix {
		newColumn := make([]string, len(matrix[0]))
		for columnIndex, seat := range column {
			newColumn[columnIndex] = seat
			if seat == floor {
				continue // nothing changes here
			}
			if seat == empty {
				if rowIndex != 0 && matrix[rowIndex-1][columnIndex] == occupied {
					continue // different seat/floor found on the left, so won't be occupying this one
				}
				if rowIndex != 0 && columnIndex != 0 && matrix[rowIndex-1][columnIndex-1] == occupied {
					continue // different seat/floor found on upper left diagonal
				}
				if columnIndex != 0 && matrix[rowIndex][columnIndex-1] == occupied {
					continue // different seat/floor found above
				}
				if columnIndex != 0 && rowIndex != len(matrix)-1 && matrix[rowIndex+1][columnIndex-1] == occupied {
					continue // different seat/floor found on upper right diagonal
				}
				if rowIndex != len(matrix)-1 && matrix[rowIndex+1][columnIndex] == occupied {
					continue // different seat/floor found on the right
				}
				if rowIndex != len(matrix)-1 && columnIndex != len(column)-1 && matrix[rowIndex+1][columnIndex+1] == occupied {
					continue // different seat/floor found on lower right diagonal
				}
				if columnIndex != len(column)-1 && matrix[rowIndex][columnIndex+1] == occupied {
					continue // different seat/floor found just below
				}
				if columnIndex != len(column)-1 && rowIndex != 0 && matrix[rowIndex-1][columnIndex+1] == occupied {
					continue // different seat/floor found on lower left diagonal
				}
				changed = true
				newColumn[columnIndex] = occupied
			} else {
				var numOfOccupied int
				if rowIndex != 0 && matrix[rowIndex-1][columnIndex] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if rowIndex != 0 && columnIndex != 0 && matrix[rowIndex-1][columnIndex-1] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if columnIndex != 0 && matrix[rowIndex][columnIndex-1] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if columnIndex != 0 && rowIndex != len(matrix)-1 && matrix[rowIndex+1][columnIndex-1] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if rowIndex != len(matrix)-1 && matrix[rowIndex+1][columnIndex] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if rowIndex != len(matrix)-1 && columnIndex != len(column)-1 && matrix[rowIndex+1][columnIndex+1] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if columnIndex != len(column)-1 && matrix[rowIndex][columnIndex+1] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
				if columnIndex != len(column)-1 && rowIndex != 0 && matrix[rowIndex-1][columnIndex+1] == occupied {
					numOfOccupied++
					if numOfOccupied >= 4 {
						changed = true
						newColumn[columnIndex] = empty
						continue
					}
				}
			}
		}
		changedSeats[rowIndex] = newColumn

	}
	return changedSeats, changed
}
