package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

const passportFieldRegexp = `([a-z]{3}):([^\s\n]+)`

var requiredFields = 7

func Four() {
	b, err := ioutil.ReadFile("data/day4.txt")
	if err != nil {
		log.Fatalf("error reading input data: %v", err)
	}
	s := split(b, "\n\n")
	passports := parsePassports(s)
	valid := 0
	for _, p := range passports {
		if p.isValid() {
			valid++
		}
	}
	fmt.Printf("valid passports: %d\n", valid)
}

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (p *passport) isValid() bool {
	valid := p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.ecl != "" && p.pid != "" && p.hcl != ""
	return valid
}

func (p *passport) setField(key, value string) {
	switch key {
	case "byr":
		p.byr = value
	case "iyr":
		p.iyr = value
	case "eyr":
		p.eyr = value
	case "hgt":
		p.hgt = value
	case "ecl":
		p.ecl = value
	case "pid":
		p.pid = value
	case "cid":
		p.cid = value
	case "hcl":
		p.hcl = value
	}
}

func parsePassports(s []string) []*passport {
	re := regexp.MustCompile(passportFieldRegexp)
	passports := make([]*passport, len(s))
	for key, val := range s {
		match := re.FindAllStringSubmatch(val, -1)
		p := &passport{}
		for _, vval := range match {
			p.setField(vval[1], vval[2])
		}
		passports[key] = p
	}
	return passports
}
