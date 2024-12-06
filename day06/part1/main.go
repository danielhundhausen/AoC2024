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
  if (sign * verticalIncrement != 1) && (sign * horizontalIncrement != 1) {
    panic("Increments are messed up")
  }
  for (*i != 0) && (*j != 0) && (*i != yDim - 1) && (*j != xDim - 1) &&
    (guardMap[*i + verticalIncrement][*j + horizontalIncrement] != "#") {
    *i += verticalIncrement
    *j += horizontalIncrement
    guardMap[*i][*j] = "X"
  }
}

func countDistinctPositions(guardMap [][]string) int {
  var n int
  for _, line := range guardMap {
    for _, char := range line {
      if char == "X" {
        n++
      }
    }
  }
  return n
}

func main() {
  // Load Map
  guardMap := loadMap()
  // Find Dimensions and Initial Position
  i, j := findInitialPosition(&guardMap)
  yDim, xDim := getDimsOfMap(&guardMap)
  // Move and count steps
  var upOrRight int = 1 // 1 = up/down; 0 = right/left
  var sign int = -1 // 1 = down/right; -1 = up/left
  var counter int = 0
  for (i != 0) && (j != 0) && (i != yDim) && (j != xDim) {
    moveToNextObstacle(xDim, yDim, guardMap, &i, &j, upOrRight, sign)
    // Check if boundry is reached
    if (i == yDim - 1) || (j == xDim - 1)  {
      break
    }
    upOrRight = (1 - upOrRight)
    if (counter % 2) == 0 {
      sign *= -1
    }
    counter++
  }
  for _, x := range guardMap {
    fmt.Println(x)
  }
  fmt.Println("Total number of Steps: ", countDistinctPositions(guardMap))
}
