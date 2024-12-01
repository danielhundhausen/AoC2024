package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readArrays() ([]int, []int) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error openning the file", err)
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

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
	var x, y []int = readArrays()

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
