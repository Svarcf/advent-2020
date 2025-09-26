package tasks

import (
	"fmt"
	"svarcf/advent-2020/common"
)

func Task3() {
	file, err := common.ReadFile("input/data3.txt")
	if err != nil {
		fmt.Println("error reading a file", err)
		panic(err)
	}

	i := 0
	j := 0
	result := 0

	lineLength := len(file[0])
	for i < len(file) {
		j += 3
		i += 1

		if i >= len(file) {
			break
		}
		if j >= lineLength {
			j = j - lineLength
		}
		if file[i][j] == '#' {
			result++
		}
	}

	fmt.Println(result)
}
