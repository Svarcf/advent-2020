package tasks

import (
	"fmt"
	"strconv"
	"strings"
	"svarcf/advent-2020/common"
)

func Task2_2() {
	file, err := common.ReadFile("input/data2.txt")
	if err != nil {
		fmt.Println("error reading a file", err)
		panic(err)
	}

	result := 0

	for _, line := range file {
		words := strings.Fields(line)
		bounds := strings.Split(words[0], "-")
		lowerChar := string(bounds[0])
		upperChar := string(bounds[1])
		targetChar := string(words[1][0])

		lower, err := strconv.Atoi(lowerChar)
		if err != nil {
			panic(err)
		}
		upper, err := strconv.Atoi(upperChar)
		if err != nil {
			panic(err)
		}
		upper = upper - 1
		lower = lower - 1

		word := string(words[2])
		if (word[lower] == targetChar[0] && word[upper] != targetChar[0]) || (word[lower] != targetChar[0] && word[upper] == targetChar[0]) {
			result++
		}

	}
	fmt.Println(result)

}
