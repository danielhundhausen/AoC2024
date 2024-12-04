package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"os"
	"regexp"
	"strings"
)

func countHF(s string) int {
	re := regexp.MustCompile(`XMAS`)
	return len(re.FindAllString(s, -1))
}

func countHB(s string) int {
	re := regexp.MustCompile(`SAMX`)
	return len(re.FindAllString(s, -1))
}

func runeToN(r rune) float64 {
	switch r {
	case 'M':
		return 1
	case 'A':
		return 2
	case 'S':
		return 3
	default:
		return 0
	}
}

func convertToMatrix(s string) mat.Dense {
	var lines []string = strings.Split(s, "\n")
	var lenLines int = len(strings.Split(lines[0], ""))

	matrix := mat.NewDense(len(lines), lenLines, nil)
	for i, l := range lines {
		if len(l) < 1 {
			continue
		}
		for j, c := range l {
			matrix.Set(i, j, runeToN(c))
		}
	}
	return *matrix
}

func nXMSAMSMatches(matrix mat.Dense) int {
	var nMatches int
	n, m := matrix.Dims()
	for i := 0; i < n-2; i++ {
		for j := 0; j < m-2; j++ {
			if matrix.At(i, j) == 1 &&
				matrix.At(i+2, j) == 3 &&
				matrix.At(i+1, j+1) == 2 &&
				matrix.At(i, j+2) == 1 &&
				matrix.At(i+2, j+2) == 3 {
				nMatches += 1
			}
		}
	}
	return nMatches
}

func nXMMASSMatches(matrix mat.Dense) int {
	var nMatches int
	n, m := matrix.Dims()
	for i := 0; i < n-2; i++ {
		for j := 0; j < m-2; j++ {
			if matrix.At(i, j) == 1 &&
				matrix.At(i+2, j) == 1 &&
				matrix.At(i+1, j+1) == 2 &&
				matrix.At(i, j+2) == 3 &&
				matrix.At(i+2, j+2) == 3 {
				nMatches += 1
			}
		}
	}
	return nMatches
}

func main() {
	// Read the entire file as one
	f, _ := os.ReadFile("day04/input")
	file_as_string := string(f)
	matrix := convertToMatrix(file_as_string)

	var nXmas int
	var nP int
	nP = nXMSAMSMatches(matrix)
	fmt.Println(nP)
	nXmas += nP
	nP = nXMMASSMatches(matrix)
	fmt.Println(nP)
	nXmas += nP
	// fmt.Println(nXmas)
	n, m := matrix.Dims()
	linspaceN := make([]int, n)
	linspaceM := make([]int, m)
	for i := 0; i < n; i++ {
		linspaceN[i] = n - i - 1
	}
	for i := 0; i < m; i++ {
		linspaceM[i] = m - i - 1
	}
	// Permute Rows
	matrix.PermuteRows(linspaceN, false)
	nP = nXMSAMSMatches(matrix)
	fmt.Println(nP)
	nXmas += nP
	matrix.PermuteRows(linspaceN, false)
	// fmt.Println(nXmas)
	// Permute Cols
	matrix.PermuteCols(linspaceM, false)
	nP = nXMMASSMatches(matrix)
	fmt.Println(nP)
	nXmas += nP
	fmt.Println("-----")
	fmt.Println(nXmas)
}
