package main

import (
	"AoC2024/day01/shared"
	"fmt"
)

func countFrequencies(numbers []int) map[int]int {
	frequencies := make(map[int]int)
	for _, num := range numbers {
		frequencies[num]++
	}
	return frequencies
}

func main() {
	var x, y []int = shared.ReadArrays()

	var freqMap map[int]int = countFrequencies(y)
	var similarityScore int
	for i := 0; i < len(x); i++ {
		similarityScore += x[i] * freqMap[x[i]]
	}
	fmt.Println(similarityScore)
}
