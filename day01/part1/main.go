package main

import (
	"AoC2024/day01/shared"
	"fmt"
	"sort"
)

func getDiff(x []int, y []int) []uint {
	var diff []uint
	for i := 0; i < len(x); i++ {
		if x[i] > y[i] {
			diff = append(diff, uint(x[i]-y[i]))
		} else {
			diff = append(diff, uint(y[i]-x[i]))
		}
	}
	return diff
}

func main() {
	var x, y []int = shared.ReadArrays()

	// Sort arrays
	sort.Ints(x)
	sort.Ints(y)
	var diff []uint = getDiff(x, y)
	var sum uint
	for i := 0; i < len(diff); i++ {
		sum += diff[i]
	}
	fmt.Println(sum)
}
