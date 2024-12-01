package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadArrays() ([]int, []int) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error openning the file", err)
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
