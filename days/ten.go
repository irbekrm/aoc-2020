package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

func Ten() {
	b, err := ioutil.ReadFile("data/day10.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	nums, err := toInts(s)
	if err != nil {
		log.Fatalf("error converting input data to integers: %v", err)
	}
	sort.Ints(nums)
	adapter := nums[len(nums)-1] + 3
	nums = append([]int{0}, nums...)
	nums = append(nums, adapter)
	fmt.Println(nums)
	product := findDiffs(nums)
	fmt.Printf("Differences of one x differences of three is %d\n", product)
}

func findDiffs(nums []int) int {
	var ones, threes int
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 3 {
			threes++
		}
		if diff == 1 {
			ones++
		}
	}
	return ones * threes
}
