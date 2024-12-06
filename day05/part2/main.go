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

func nSortedFailed(update []int, rules [][2]int) (bool, int) {
  var updateReSorted  bool
  // Brute force way to "fix" rulebreaks caused by sccessive rule application
  for i := 0; i < 5; i++ { 
    // Check every rule
    for _, rule := range rules {
      var idxFirst int = -1
      var idxSecond int = -1
      for i, x := range update {
        // Find idx of first and second rule entry
        if x == rule[0] {
          idxFirst = i
        } else if x == rule[1] {
          idxSecond = i
        }
        if (idxFirst > -1) && (idxSecond > -1) {
          if idxFirst > idxSecond {
            update[idxFirst], update[idxSecond] = update[idxSecond], update[idxFirst]
            updateReSorted = true
            break
          } else {
            break
          }
        }
      }
    }
  }
  if updateReSorted {
    return true, update[len(update) / 2]
  }
  return false, 0
}

func main() {
	rules, updates := loadRulesAndUpdates()
  fmt.Println("N Rule: ", len(rules))
  fmt.Println("N Updates: ", len(updates))
  fmt.Println("----------------")
  var sumUpdates int
  var n int
  for i := 0; i < len(updates); i++ {
    _, n = nSortedFailed(updates[i], rules)
    sumUpdates += n
  }
	fmt.Println(sumUpdates)
}
