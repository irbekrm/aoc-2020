package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
)

func Thirteen() {
	b, err := ioutil.ReadFile("data/day13.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	num, err := findBus(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}

func findBus(s []string) (int, error) {
	t, err := strconv.Atoi(s[0])
	if err != nil {
		return -1, err
	}
	re := regexp.MustCompile(`\d+`)
	busData := re.FindAllString(s[1], -1)
	shortest := math.Inf(1)
	buses := make(map[int]int, len(busData))
	for _, val := range busData {
		id, err := strconv.Atoi(val)
		if err != nil {
			return -1, err
		}
		wait := id - t%id
		buses[wait] = id
		if w := float64(wait); w < shortest {
			shortest = w
		}

	}
	return buses[int(shortest)] * int(shortest), nil
}
