package main

import (
  "fmt"
)

func main() {
  x := make([]int, 77)
  x = append(x, 44)
  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(cap(x))

  var y []int
  // y = [2]int{1, 2}
  fmt.Println(y)
  fmt.Println(len(y))
  fmt.Println(cap(y))
}
