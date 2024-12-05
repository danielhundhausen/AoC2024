package main

import (
	"os"
	"bufio"
  "regexp"
  "strconv"
	"fmt"
)

func loadRulesAndUpdates() ([][2]int, [][]int) {
	file, err := os.Open("day05/input")
	if err != nil {
		panic("File opening failed!")
	}
	defer file.Close()

	var rules [][2]int
	var updates [][]int
	var part2 bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) < 1 {
			part2 = true
			continue
		}
		if !part2 {
			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllString(line, -1)
			if len(matches) != 2 {
		    panic("Expected rules of the form `int|int`!")
			}
			var nextRule [2]int
			for i, n := range matches {
				nextRule[i], _ = strconv.Atoi(n)
			}
			rules = append(rules, nextRule)
		} else {
			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllString(line, -1)
			var nextEntry []int
			for _, n := range matches {
        num, _ := strconv.Atoi(n)
				nextEntry = append(nextEntry, num)
			}
			updates = append(updates, nextEntry)
		}
	}
	return rules, updates
}

func main() {
	rules, updates := loadRulesAndUpdates()
	fmt.Println(rules)
	fmt.Println(updates)
}
