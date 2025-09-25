package tasks

import (
	"fmt"
	"sort"
	"strconv"
	"svarcf/advent-2020/common"
)

func Task1_2() {
	file, err := common.ReadFile("input/data1.txt")
	if err != nil {
		panic(err)
	}
	arr := make([]int, len(file))

	for index, line := range file {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		arr[index] = i
	}
	sort.Ints(arr)

	for i := range len(arr) - 2 {
		for j := i + 1; j < len(arr)-1; j++ {
			temp := 2020 - arr[i] - arr[j]

			for k := j + 1; k < len(arr); k++ {
				if arr[k] == temp {
					fmt.Println(arr[k] * arr[i] * arr[j])
					return
				}
				if arr[k] > temp {
					break
				}
			}
		}
	}
}
