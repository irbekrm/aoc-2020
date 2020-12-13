package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

const passportFieldRegexp = `([a-z]{3}):([^\s\n]+)`

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
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid passportField
}

type passportField interface {
	isValid() bool
}

type byr string

func (b byr) isValid() bool {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return false
	}
	return i >= 1920 && i <= 2002
}

type iyr string

func (i iyr) isValid() bool {
	num, err := strconv.Atoi(string(i))
	if err != nil {
		return false
	}
	return num >= 2010 && num <= 2020
}

type eyr string

func (e eyr) isValid() bool {
	num, err := strconv.Atoi(string(e))
	if err != nil {
		return false
	}
	return num >= 2020 && num <= 2030
}

type hgt string

func (h hgt) isValid() bool {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	if !re.MatchString(string(h)) {
		return false
	}
	match := re.FindAllStringSubmatch(string(h), -1)
	num, err := strconv.Atoi(match[0][1])
	if err != nil {
		return false
	}
	if match[0][2] == "cm" {
		return num >= 150 && num <= 193
	}
	if match[0][2] == "in" {
		return num >= 59 && num <= 76
	}
	return false
}

type hcl string

func (h hcl) isValid() bool {
	re := regexp.MustCompile(`^#[0-9a-z]{6}`)
	return re.MatchString(string(h))
}

type ecl string

func (e ecl) isValid() bool {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return re.MatchString(string(e))
}

type pid string

func (p pid) isValid() bool {
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(string(p))
}

func (p *passport) hasCompulsoryFields() bool {
	valid := p.byr != nil && p.iyr != nil && p.eyr != nil && p.hgt != nil && p.ecl != nil && p.pid != nil && p.hcl != nil
	return valid
}

func (p *passport) isValid() bool {
	if !p.hasCompulsoryFields() {
		return false
	}
	return p.byr.isValid() && p.iyr.isValid() && p.eyr.isValid() && p.hgt.isValid() && p.ecl.isValid() && p.pid.isValid() && p.hcl.isValid()
}

func (p *passport) setField(key, value string) {
	switch key {
	case "byr":
		p.byr = byr(value)
	case "iyr":
		p.iyr = iyr(value)
	case "eyr":
		p.eyr = eyr(value)
	case "hgt":
		p.hgt = hgt(value)
	case "ecl":
		p.ecl = ecl(value)
	case "pid":
		p.pid = pid(value)
	case "hcl":
		p.hcl = hcl(value)
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
