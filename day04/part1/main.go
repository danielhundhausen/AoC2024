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
	case 'X':
		return 1
	case 'M':
		return 2
	case 'A':
		return 3
	case 'S':
		return 4
	default:
		return 0
	}
}

func nToRune(n float64) rune {
	switch n {
	case 1:
		return 'X'
	case 2:
		return 'M'
	case 3:
		return 'A'
	case 4:
		return 'S'
	default:
		return 'Z'
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

func convertToString(matrix mat.Matrix) string {
	var s string
	n, m := matrix.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s += string(nToRune(matrix.At(i, j)))
		}
		s += "\n"
	}
	return s
}

func nDiagMatchesSlash(matrix mat.Dense) int {
	var nMatches int
	n, m := matrix.Dims()
	for i := 0; i < n-3; i++ {
		for j := 0; j < m-3; j++ {
			if matrix.At(i, j) == 1 &&
				matrix.At(i+1, j+1) == 2 &&
				matrix.At(i+2, j+2) == 3 &&
				matrix.At(i+3, j+3) == 4 {
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

	var n_xmas int
	// Cases:
	// Handle with normal regex
	// horizontal forward
	// horizontal backward
	n_xmas += countHF(file_as_string)
	n_xmas += countHB(file_as_string)

	// Transopse matrix and handle normally
	// vertical forward
	// vertical backward
	matrix := convertToMatrix(file_as_string)
	var file_as_string_T string = convertToString(matrix.T())
	n_xmas += countHF(file_as_string_T)
	n_xmas += countHB(file_as_string_T)

	// Convolution like scan
	// diagnoal forward down
	n_xmas += nDiagMatchesSlash(matrix)
	n, m := matrix.Dims()
	linspaceN := make([]int, n)
	linspaceM := make([]int, m)
	for i := 0; i < n; i++ {
		linspaceN[i] = n - i - 1
	}
	for i := 0; i < m; i++ {
		linspaceM[i] = m - i - 1
	}
	// diagonal backward down
	matrix.PermuteCols(linspaceM, false)
	n_xmas += nDiagMatchesSlash(matrix)
	// diagonal backward up
	matrix.PermuteRows(linspaceN, false)
	n_xmas += nDiagMatchesSlash(matrix)
	// diagonal forward up
	matrix.PermuteCols(linspaceM, false)
	n_xmas += nDiagMatchesSlash(matrix)
	fmt.Println(n_xmas)
}
