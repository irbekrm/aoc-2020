package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

const requiredPasswdRegexp = `^(\d+)-(\d+) ([a-z]): ([a-z]+)$`

func Two() {
	b, err := ioutil.ReadFile("data/day2.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n")
	passwds := make([]*passwd, len(s))
	for k, v := range s {
		passwd, err := stringToPasswd(v)
		if err != nil {
			log.Fatalf("error parsing data: %v", err)
		}
		passwds[k] = passwd
	}
	valid := 0
	for _, v := range passwds {
		if v.isValidTobbogan() {
			valid++
		}
	}
	fmt.Printf("The number of valid passwords: %d\n", valid)
}

type passwd struct {
	req      rune
	min, max int
	content  []rune
}

func (p *passwd) isValidTobbogan() bool {
	return p.content[p.min-1] == p.req && !(p.content[p.max-1] == p.req) || !(p.content[p.min-1] == p.req) && (p.content[p.max-1] == p.req)
}

func (p *passwd) isValidSled() bool {
	if len(p.content) < p.min {
		return false
	}
	got := 0
	for k, v := range p.content {
		if v == p.req {
			got++
		}
		if got > p.max {
			return false
		}
		// there are less runes remaining to parse than the number of reequired letters still to be found
		if len(p.content)-k-1 < (p.min - got) {
			return false
		}
	}
	return got >= p.min && got <= p.max
}

func stringToPasswd(s string) (*passwd, error) {
	re := regexp.MustCompile(requiredPasswdRegexp)
	m := re.FindAllStringSubmatch(s, -1)
	if len(m) < 1 {
		return nil, errors.New("empty match")
	}
	matches := m[0]
	if len(matches) < 5 {
		return nil, fmt.Errorf("not enough matches in %s, expected 5 got %d", s, len(matches))
	}
	var matchErr error
	min, err := strconv.Atoi(matches[1])
	if err != nil {
		matchErr = errors.Wrap(matchErr, fmt.Sprintf("error converting %v to int: %v", matches[1], err))
	}
	max, err := strconv.Atoi(matches[2])
	if err != nil {
		matchErr = errors.Wrap(matchErr, fmt.Sprintf("error converting %v to int: %v", matches[2], err))
	}
	req := []rune(matches[3])[0]
	content := []rune(matches[4])
	if matchErr != nil {
		return nil, matchErr
	}
	return &passwd{
		min:     min,
		max:     max,
		req:     req,
		content: content,
	}, nil
}
