package main

import "fmt"

// Notes:
// After thinking a lot of how to implement this, I came to the conclusion
// that the best way to complete this in Go is to duplicate code...
// I was either duplicate methods or pass both Stacks into the method.
// This way it's type checking the Stack pass through.
// Another implementation could have been using a struct that held both
// EvenStack & OddStack, which then would assign and grab accordingly.......
// example:
//
// type Stack struct {
//   EventStack []int
//   OddStack []int
// }

// Declaration of Stack interface
type Stack interface {
  Push(int)
  Pop()
  Get()
  GetNums()
}

// Declaration of EvenStack/OddStack, using an Integer Slice
// Slices are more powerful arrays
type EvenStack []int
type OddStack []int

// Defines Push methods, implementing LIFO
func (s *EvenStack) Push(item int) { // EvenStack type only
  // verifies item is greater than 0
  if item < 0 {
    // Prints out error, re-formated to slice specification `%v`
    fmt.Print(fmt.Errorf("item is not positive: %v", item))
  } else if num := item%2; num != 0 { // verifies item is even
    fmt.Print(fmt.Errorf("item is not even: %v", item))
  } else {
    *s  = append(*s, item) // appends item to Slice(which is a pointer)
  }
}

func (s *OddStack) Push(item int) { // OddStack type only
  // verifies item is greater than 0
  if item < 0 {
    // Prints out error, re-formated to slice specification `%v`
    fmt.Print(fmt.Errorf("item is not positive: %v", item))
  } else if num := item%2; num == 0 { // verifies item is odd
    fmt.Print(fmt.Errorf("item is not odd: %v", item))
  } else {
    *s = append(*s, item) // appends item to Slice(which is a pointer)
  }
}

// Defines Pop methods, implementing LIFO
func (s *EvenStack) Pop() {
  c := *s // copy slice
  c = append(c[1:]) // potential memory leak without nil'ing out first item
  *s = c
}

func (s *OddStack) Pop() {
  c := *s
  c = append(c[1:])
  *s = c
}

// Defines Get methods
func (s EvenStack) Get() int {
  return s[0] // returns first item in slice
}

func (s OddStack) Get() int {
  return s[0] // returns first item in slice
}

// Defines GetNum methods
func (s EvenStack) GetNums() int {
  return len(s) // returns length of slice
}

func (s OddStack) GetNums() int {
  return len(s) // returns length of slice
}
