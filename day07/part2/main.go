package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func loadEquations() map[int][]int {
	file, _ := os.Open("day07/input")
	var eqns map[int][]int = make(map[int][]int)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		var result int
		for i, n := range re.FindAllString(line, -1) {
			_n, _ := strconv.Atoi(n)
			if i == 0 {
				result = _n
				continue
			}
			numbers = append(numbers, _n)
		}
		eqns[result] = numbers
	}
	return eqns
}

func addOrMultiply(a, b int, add int) int {
	if add == 0 {
		return a + b
	} else if add == 1 {
    v, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
    return v
  }
	return a * b
}

func powInt(x, y int) int {
	if y == 0 {
		return 1
	}
	return int(math.Pow(float64(x), float64(y)))
}

func getOperatorPermutations(n int) [][]int {
    var perms [][]int
    totalPerms := pow(3, n) // 3^n permutations for three possible values
    
    for i := 0; i < totalPerms; i++ {
        perm := make([]int, n)
        num := i
        
        // Convert to base-3 representation
        for j := 0; j < n; j++ {
            perm[j] = num % 3  // Get remainder when divided by 3
            num /= 3           // Integer division by 3 for next digit
        }
        perms = append(perms, perm)
    }
    return perms
}

// Helper function to calculate power since math.Pow returns float64
func pow(base, exp int) int {
    result := 1
    for i := 0; i < exp; i++ {
        result *= base
    }
    return result
}

func checkEqn(resultTruth int, operands []int) int {
	var result int
	operators := getOperatorPermutations(len(operands) - 1)
	for _, operatorPerm := range operators {
    result = operands[0]
		for i := 1; i < len(operands); i++ {
			result = addOrMultiply(result, operands[i], operatorPerm[i-1])
		}
		if result == resultTruth {
			return result
		}
	}
	return 0
}

func main() {
	eqns := loadEquations()
	var sumCorrect int
	for result, operands := range eqns {
		fmt.Println(result, operands)
		sumCorrect += checkEqn(result, operands)
	}
  fmt.Println("Final Sum: ", sumCorrect)
}
