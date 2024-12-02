package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func loadInputToArrays(fpath string) [][]int {
	// Open file and schedule closing it when funciton returns
	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arrays [][]int
	var _arr []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_arr = []int{}
		nums := strings.Fields(scanner.Text())
		for i := 0; i < len(nums); i++ {
			number, _ := strconv.Atoi(nums[i])
			if number < 1 {
				fmt.Println(nums)
			}
			_arr = append(_arr, number)
		}
		arrays = append(arrays, _arr)
	}

	return arrays
}

func checkProblemDamper(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		// Fill the dampedArr
		var dampedArr []int
		for k := 0; k < len(arr); k++ {
			if k == i {
				continue
			}
			dampedArr = append(dampedArr, arr[k])
		}

		var diffs []int = make([]int, len(dampedArr)-1)
		for i := 0; i < len(dampedArr)-1; i++ {
			diffs[i] = dampedArr[i+1] - dampedArr[i]
		}

		maxDiff, minDiff := slices.Max(diffs), slices.Min(diffs)
		if ((minDiff >= 1) && (maxDiff <= 3))|| ((maxDiff <= -1) && (minDiff >= -3)) {
			return true
		}
	}
	return false
}

func isAscOrDescending(arr []int) bool {
	var diffs []int = make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		diffs[i] = arr[i+1] - arr[i]
	}

	maxDiff, minDiff := slices.Max(diffs), slices.Min(diffs)
	return ((minDiff >= 1) && (maxDiff <= 3))|| ((maxDiff <= -1) && (minDiff >= -3))
}

func main() {
	var arrays [][]int = loadInputToArrays("day02/input")
	var safeReports int
	for _, report := range arrays {
		if !isAscOrDescending(report) && !checkProblemDamper(report) {
			continue
		}
		safeReports += 1
	}
	fmt.Println(safeReports)
}
