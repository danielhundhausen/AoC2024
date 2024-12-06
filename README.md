# AoC 2024

Language: Go

# Learning Notes
- Day 1:
  - Go installation, go.mod file
- Day 4:
  - There is no `np.arange/linspace` equivalent
  - Instead `make([]int, N)` and a loop to fill the slice
- Day 5:
  - `Stringer` interface is the equivalent of `__srt__` in Python
  - No real inheritance, just implementation of interfaces
- Day 6:
  - `func (t *test) Test() bool` to attach a function Test to the test struct. Need to pass pointer with `*` not to copy original object.
  - for loops can be over
    - `{ }`: implicitly true
    - `bool`: execute if true
    - `i := 0; i < X; i++`: Standard 3 expression ;seped
    - `_, x := range <...>`: range + one of {array, slice, string, map, channel}
  - Block scope: Variables declared inside a loop don't overwrite those decleared outside
  - No `while` loops, implemented by `for <condition>`
