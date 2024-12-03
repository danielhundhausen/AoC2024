// Package main solves day03:part1 of AoC2024.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// readMemoryInput reads the contents of `day03/input` and returns them as a string.
func readMemoryInput() string {
	body, _ := os.ReadFile("day03/input")
	return string(body)
}

// filterCorrectStatements takes a raw memory string as input and retruns an array of
// valid `mul(a,b)` statements.
func filterCorrectStatements(mem string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return re.FindAllString(mem, -1)
}

// main computes the sum of all valid `mul(a,b)` statements
func main() {
	mem := readMemoryInput()
	mulCombinations := filterCorrectStatements(mem)
	var sum int
	re := regexp.MustCompile(`\d+`)
	for _, x := range mulCombinations {
		nums := re.FindAllString(x, 2)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		sum += a * b
	}
	fmt.Println(sum)
}
