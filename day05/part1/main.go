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

func updatePassesRules(update []int, rules [][2]int) (bool, int) {

  for _, rule := range rules {
    var idxFirst int = -1
    var idxSecond int = -1
    for i, x := range update {
      if x == rule[0] {
        idxFirst = i
      } else if x == rule[1] {
        idxSecond = i
      }
      if (idxFirst > -1) && (idxSecond > -1) {
        break
      }
    }
    if idxFirst > idxSecond && idxSecond > -1 {
      return false, 0
    }
  }

  var idxMiddle int = int(len(update) / 2)
  return true, update[idxMiddle]
}

func main() {
	rules, updates := loadRulesAndUpdates()
  var sumUpdates int
  var n int
  for i := 0; i < len(updates); i++ {
    _, n = updatePassesRules(updates[i], rules)
    sumUpdates += n
  }
	fmt.Println(sumUpdates)
}
