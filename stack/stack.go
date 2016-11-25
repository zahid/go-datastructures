package stack

/*
Stack is an data structure which is a linear collection of elements with two primary methods, push,
which adds an element to the collection, and pop, which removes the most recently added element.
The elements are added and removed in LIFO (last in, first out).
*/

/*
This is an interface representing a Stack
*/
type Stack interface {
	Push()
	Pop()
	Peek()
	IsEmpty()
	Clear()
}

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
SliceStack is an implementation of a Stack based on a slice.
*/
type SliceStack struct {
	elements Elements
}

/*
Push inserts an element into the Stack
*/
func (s *SliceStack) Push(e Element) {
	s.elements = append(s.elements, e)
}

/*
Pop returns the element of the last inserted element. The top element is also removed from the Stack
*/
func (s *SliceStack) Pop() Element {
	if len(s.elements) == 0 {
		return nil
	}
	top := s.elements[0]
	s.elements = s.elements[1:]
	return top
}

/*
Peek returns the element of the last inserted element. This does not remove the element from the Stack
*/
func (s *SliceStack) Peek() Element {
	if len(s.elements) == 0 {
		return nil
	}
	top := s.elements[0]
	return top
}

/*
IsEmpty returns a boolean indicating whether or not the Stack is empty
*/
func (s *SliceStack) IsEmpty() bool {
	return len(s.elements) == 0
}

/*
Clear empties the Stack
*/
func (s *SliceStack) Clear() {
	s.elements = Elements{}
}
