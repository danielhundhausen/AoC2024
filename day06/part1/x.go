package main

import (
  "fmt"
)

type test struct {
  i int
}

func (t *test) Test() bool {
  fmt.Println(t.i)
  if t.i < 3 {
    t.i = 4
    return true
  }
  return false
}

func main() {
  var tObj test
  for tObj.Test() {
    fmt.Println("foo")
  }
}
