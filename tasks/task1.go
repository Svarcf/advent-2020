package tasks

import (
	"fmt"
	"strconv"
	"svarcf/advent-2020/common"
)

func Task1() {
	file, err := common.ReadFile("input/data1.txt")
	if err != nil {
		fmt.Println("error reading a file", err)
		panic(err)
	}

	goal := 2020
	valueMap := make(map[int]bool)
	for _, line := range file {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		valueMap[i] = true
	}

	for _, line := range file {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		deduction := goal - i
		_, ok := valueMap[deduction]
		if ok {
			fmt.Println(deduction * i)
			return
		}
	}
}
