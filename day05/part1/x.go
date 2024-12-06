package main

import (
  "fmt"
)

func main() {
  var x []int = []int{1, 2, 3, 4, 5}
  fmt.Println(len(x))
  fmt.Println(len(x) / 2)
  fmt.Println(x[len(x) / 2])
}
