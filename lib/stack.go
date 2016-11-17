package main

import (
	"fmt"
)

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
	fmt.Printf("Pushing (%v, %T)\n", e, e)
	s.elements = append(s.elements, e)
}

func (s *Stack) pop() {
	// something's funny about this index
	lastIndex := len(s.elements) - 1
	s.elements = s.elements[1:lastIndex]
}

func main() {
	stackOfPapers := Stack{}

	stackOfPapers.push("bill")
	stackOfPapers.push("reciept")
	stackOfPapers.push("menu")
	stackOfPapers.push("picture")
	stackOfPapers.pop()

	fmt.Printf("(%v, %T)\n", stackOfPapers, stackOfPapers)
}
