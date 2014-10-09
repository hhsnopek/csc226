package assignment5

// Notes:
// After thinking a lot of how to implement this, I first came to the conclusion
// that the best way to complete this in Go is to duplicate code...
// It was either duplicate methods or pass both Stacks into the method.
// This way it's type checking the Stack pass through.
// My final implementation was using a struct that held both EvenStack &
// OddStack

// Syntax explainations
// `:=` is short hand for delcaration and initialization of a var within a
// function, `:=` is not available outside a function.
// `*` creates a pointer, allowing us to pass by reference instead of value

// Function Declaration
// Basic fn:
// func Name() ReturnType { ... }
// Basic Method:
// func (s ObjectType) Name() ReturnType { ... }

// Declaration of Stack interface
// This isn't nessecary considering I did not use Stack within a method
// Rather than using GetEven or GetOdds, I could have used a single method
// that would take Stack, for generic use.
type Stacker interface {
	Push(int)
	Pop()
	Get()
	GetNums()
}

// Declaration of StackObject struct, using an Integer Slice
// Slices are more powerful arrays
type EvenStack []int
type OddStack []int
type StackObject struct {
	EvenStack []int
	OddStack  []int
}

// Defines Push methods, implementing LIFO
// append usage below is interesting, as this only works with the `...`
// without it, append only works when you append existing data to a slice.
// ex: append(slice, data)
func (s *StackObject) Push(item int) {
	// verifies item is greater than 0
	if item < 0 {
		// Prints out error, re-formated to slice specification `%v`
		return
	} else if (item % 2) != 0 { // verifies item is even
		c := s.EvenStack
		c = append([]int{item}, s.EvenStack...)
		s.EvenStack = c
	} else if (item % 2) == 0 { // verifies item is odd
		c := s.OddStack
		c = append([]int{item}, s.OddStack...)
		s.OddStack = c
	}
}

// Defines Pop methods, implementing LIFO
func (s *StackObject) PopEven() {
	c := *s                               // copy slice
	c.EvenStack = append(c.EvenStack[1:]) // potential memory leak
	*s = c
}

func (s *StackObject) PopOdd() {
	c := *s                             // copy slice
	c.OddStack = append(c.OddStack[1:]) // potential memory leak
	*s = c
}

// Defines Get methods
// Within Go programs it's convention with getters to not have `Get` as part of
// the method. see: http://golang.org/doc/effective_go.html#Getters
func (s StackObject) Even() int {
	return s.EvenStack[0] // returns first item in slice
}

func (s StackObject) Odd() int {
	return s.OddStack[0] // returns first item in slice
}

// Defines GetNum methods
func (s StackObject) NumsEven() int {
	return len(s.EvenStack) // returns length of slice
}

func (s StackObject) NumsOdd() int {
	return len(s.OddStack) // returns length of slice
}
