package stack

/*
We are implementing a First-In-First-Out stack
*/

/*
We have a null interface which we use as a type for an "element"
*/
type Element interface{}

/*
A slice of Elements
*/
type Elements []Element

/*
For now, this is a string type struct
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
