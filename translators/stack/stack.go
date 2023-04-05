package stack

// Stack represents generic LIFO collection of elements
type Stack[T any] []T

// Pop returns the top element by removing it from the stack
func (s *Stack[T]) Pop() T {
	if s == nil || s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	f := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return f
}

// Push adds new element to the top of the stack
func (s *Stack[T]) Push(element T) {
	if s == nil {
		return
	}
	*s = append(*s, element)
}

// IsEmpty returns true if stack has no elements
func (s *Stack[T]) IsEmpty() bool {
	return s == nil || len(*s) == 0
}
