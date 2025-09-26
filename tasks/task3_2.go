package tasks

import (
	"fmt"
	"svarcf/advent-2020/common"
)

func Task3_2() {
	file, err := common.ReadFile("input/data3.txt")
	if err != nil {
		fmt.Println("error reading a file", err)
		panic(err)
	}

	var results []int
	movements := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	lineLength := len(file[0])

	for _, move := range movements {
		result := 0
		i := 0
		j := 0
		for i < len(file) {
			j += move[0]
			i += move[1]

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
		results = append(results, result)
	}

	finalResult := 1
	for _, res := range results {
		finalResult *= res
	}
	fmt.Println(finalResult)

}
