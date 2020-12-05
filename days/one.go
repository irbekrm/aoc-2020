package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func One() {
	b, err := ioutil.ReadFile("data/day1.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b)
	ints, err := toInts(s)
	if err != nil {
		log.Fatalf("error converting input into integers: %v", err)
	}
	sort.Ints(ints)
	found, int1, int2, int3 := trippleSum2020(ints)
	if !found {
		log.Fatal("numbers not found!")
	}
	product := int1 * int2 * int3
	fmt.Printf("The two numbers are: %d %d and %d, the product is %d\n", int1, int2, int3, product)
}

func sum2020(ints []int) (bool, int, int) {
	for index, int := range ints {
		reverseIndex := len(ints) - 1
		for reverseIndex > index {
			sum := int + ints[reverseIndex]
			if sum == 2020 {
				return true, int, ints[reverseIndex]
			}
			if sum < 2020 {
				break
			}
			reverseIndex--
		}
	}
	return false, 0, 0
}

func trippleSum2020(ints []int) (bool, int, int, int) {
	for lastIndex := (len(ints) - 1); lastIndex > 1; lastIndex-- {
		firstIndex := 0
		for firstIndex < lastIndex-1 {
			intermediateSum := ints[firstIndex] + ints[lastIndex]
			if intermediateSum > 2020 {
				break
			}
			midIndex := firstIndex + 1
			for midIndex < lastIndex {
				sum := intermediateSum + ints[midIndex]
				if sum == 2020 {
					return true, ints[firstIndex], ints[midIndex], ints[lastIndex]
				}
				if sum > 2020 {
					break
				}
				midIndex++
			}
			firstIndex++
		}
	}
	return false, 0, 0, 0
}

func split(b []byte) []string {
	return strings.Split(string(b), "\n")
}

func toInts(s []string) ([]int, error) {
	ints := make([]int, len(s))
	for k, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		ints[k] = i
	}
	return ints, nil
}
