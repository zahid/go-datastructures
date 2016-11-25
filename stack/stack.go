package stack

/*
We are implementing a First-In-Last-Out stack
*/

/*
Element is an interface for an item in the Stack
*/
type Element interface{}

/*
Elements is a collection of Element structs.
It is represented using a slice
*/
type Elements []Element

/*
Stack is an data structure which is a linear collection of elements with two primary methods, push,
which adds an element to the collection, and pop, which removes the most recently added element.
The elements are added and removed in LIFO (last in, first out).
*/
type Stack struct {
	elements Elements
}

func (s *Stack) push(e Element) {
	s.elements = append(s.elements, e)
}

func (s *Stack) pop() Element {
	top := s.elements[0]
	s.elements = s.elements[1:]
	return top
}

func (s *Stack) peek() Element {
	top := s.elements[0]
	return top
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) clear() {
	s.elements = Elements{}
}
