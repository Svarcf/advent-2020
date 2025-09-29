package tasks

import (
	"fmt"
	"strings"
	"svarcf/advent-2020/common"
)

var rules = map[string]struct{}{
	"byr": {},
	"iyr": {},
	"eyr": {},
	"hgt": {},
	"hcl": {},
	"ecl": {},
	"pid": {},
}

func Task4() {
	file, err := common.ReadFileTask4("input/data4.txt")
	if err != nil {
		fmt.Println("error reading a file", err)
		panic(err)
	}

	valid := 0
	for _, line := range file {
		if isValid(line) {
			valid++
		}
	}
	fmt.Println(valid)
}

func isValid(line string) bool {
	lineSplitted := strings.Split(line, " ")
	sum := 0
	for _, props := range lineSplitted {
		prop := strings.Split(props, ":")[0]
		if _, exists := rules[prop]; exists {
			sum++
		}
	}
	test := sum == len(rules)
	return test
}
