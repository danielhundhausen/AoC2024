package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadMap() [][]string {
	file, err := os.Open("day06/input")
	if err != nil {
		panic("File opening failed!")
	}
	defer file.Close()
	var guardMap [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		guardMap = append(guardMap, chars)
	}
	return guardMap
}

func findInitialPosition(guardMap *[][]string) (i, j int) {
	for i, line := range *guardMap {
		for j, char := range line {
			if char == "^" {
				return i, j
				(*guardMap)[i][j] = "X"
			}
		}
	}
  panic("initial coordinates not found")
	return -1, -1
}

func getDimsOfMap(guardMap *[][]string) (int, int) {
	yDim := len(*guardMap)
	xDim := len((*guardMap)[0])
	return yDim, xDim
}

func moveToNextObstacle(
	xDim int,
	yDim int,
	guardMap [][]string,
	i *int,
	j *int,
	direction int,
	sign int) {
	var verticalIncrement int = sign * direction
	var horizontalIncrement int = sign * (1 - direction)
	// Assert that one of the increments is always 1
	if (sign*verticalIncrement != 1) && (sign*horizontalIncrement != 1) {
		panic("Increments are messed up")
	}
	for (*i != 0) && (*j != 0) && (*i != yDim-1) && (*j != xDim-1) &&
		(guardMap[*i+verticalIncrement][*j+horizontalIncrement] != "#") {
		*i += verticalIncrement
		*j += horizontalIncrement
		guardMap[*i][*j] = "X"
	}
}

func getDistinctPositions(guardMap [][]string) [][2]int {
	var distinctPositions [][2]int
	for i, line := range guardMap {
		for j, char := range line {
			if char == "X" {
				distinctPositions = append(distinctPositions, [2]int{i, j})
			}
		}
	}
	return distinctPositions
}

func moveThroughMap(guardMap *[][]string, iInit int, jInit int, yDim int, xDim int) bool {
	var upOrRight int = 1 // 1 = up/down; 0 = right/left
	var sign int = -1     // 1 = down/right; -1 = up/left
	var counter int = 0
  i, j := iInit, jInit
  var nMovesAllowed int = 500
  for (i != 0) && (j != 0) && (i != yDim-1) && (j != xDim-1) {
    moveToNextObstacle(xDim, yDim, *guardMap, &i, &j, upOrRight, sign)
    // Change direction
    upOrRight = (1 - upOrRight)
    if (counter % 2) == 0 {
      sign *= -1
    }
    counter++
    // Counter for when stuck
    nMovesAllowed -= 1
    if nMovesAllowed == 0 {
      return true
    }
  }
  return false
}

func main() {
	// Load Map
	guardMap := loadMap()
	// Find Dimensions and Initial Position
	iInit, jInit := findInitialPosition(&guardMap)
	yDim, xDim := getDimsOfMap(&guardMap)
	// Get the uninhibited path
	moveThroughMap(&guardMap, iInit, jInit, yDim, xDim)
	pathPositions := getDistinctPositions(guardMap)
  fmt.Println("N Path Positions: ", len(pathPositions))
	// Move and count steps
	var nLoopBlocks int
	for _, x := range pathPositions {
    if (x[0] == iInit) && (x[1] == jInit) {
      continue
    }
		guardMap = loadMap()
		guardMap[x[0]][x[1]] = "#"
		if moveThroughMap(&guardMap, iInit, jInit, yDim, xDim) {
			nLoopBlocks++
		}
	}
	fmt.Println("Total number of Looping Positions: ", nLoopBlocks)
}
