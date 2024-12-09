package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadMap() [][]string {
	file, err := os.Open("day08/input")
	if err != nil {
		panic("File opening failed!")
	}
	defer file.Close()
	var antennaMap [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		antennaMap = append(antennaMap, chars)
	}
	return antennaMap
}

func getDimsOfMap(antennaMap *[][]string) (int, int) {
	yDim := len(*antennaMap)
	xDim := len((*antennaMap)[0])
	return yDim, xDim
}

func findDistinctAntennas(antennaMap [][]string) []string {
  antennas := make([]string, 1)
  for i := 0; i < len(antennaMap); i++{
    for j := 0; j < len(antennaMap[i]); j++ {
      if (antennaMap[i][j] != ".") && (antennaMap[i][j] != ""){
        antennas = append(antennas, antennaMap[i][j])
      }
    }
  }
  return antennas
}

func findCoordinatesOfAntenna(antennaMap [][]string, antenna string) [][2]int {
	var distinctPositions [][2]int
	for i, line := range antennaMap {
		for j, char := range line {
			if char == antenna {
				distinctPositions = append(distinctPositions, [2]int{i, j})
			}
		}
	}
	return distinctPositions
}

func findCoordinatesOfAntennas(antennaMap [][]string, antennas []string) map[string][][2]int {
  coordinateMap := make(map[string][][2]int)
  for _, antenna := range antennas {
    coordinateMap[antenna] = findCoordinatesOfAntenna(antennaMap, antenna)
  }
  return coordinateMap
}

func main() {
	// Load Map
	antennaMap := loadMap()
	yDim, xDim := getDimsOfMap(&antennaMap)
  fmt.Println("Dims of map: ", xDim, yDim)

	// Find The types of antennas there are
  distinctAntennas := findDistinctAntennas(antennaMap)
  fmt.Println("")
  fmt.Println("Distinct Antennas")
  fmt.Println("-----------------")
  fmt.Println(distinctAntennas)
  fmt.Println("")
  fmt.Println("Coordinates")
  fmt.Println("-----------")
  maps := findCoordinatesOfAntennas(antennaMap, distinctAntennas)
  for k, v := range maps {
 	  fmt.Println(k, v)
  }
}
