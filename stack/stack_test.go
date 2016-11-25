package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stackOfPapers := SliceStack{}
	item := "Phone number"
	stackOfPapers.Push(item)

	expectedSize := 1
	actualSize := len(stackOfPapers.elements)

	if expectedSize != actualSize {
		t.Errorf("Expected size of %v, but got %v", expectedSize, actualSize)
	}

	if stackOfPapers.elements[0] != item {
		t.Errorf("Expected the element to be %v, but got %v", item, stackOfPapers.elements[0])
	}
}

func TestPop(t *testing.T) {
	stackOfPapers := SliceStack{}
	poppedNothing := stackOfPapers.Pop()

	if poppedNothing != nil {
		t.Errorf("Popping on an empty Stack should return nil, instead returned %v", poppedNothing)
	}

	item := "Dinner check"
	stackOfPapers.Push(item)

	poppedItem := stackOfPapers.Pop()

	// Check that poppedItem == item
	if item != poppedItem {
		t.Errorf("Item %v and popped item %v do not match", item, poppedItem)
	}
	// Check that the stack is empty
	if len(stackOfPapers.elements) != 0 {
		t.Errorf("Stack should be empty")
	}
}

func TestPeek(t *testing.T) {
	stackOfPapers := SliceStack{}

	peekedNothing := stackOfPapers.Peek()

	if peekedNothing != nil {
		t.Errorf("Peeking on an empty Stack should return nil, instead returned %v", peekedNothing)
	}

	item := "Dry cleaning bill"
	stackOfPapers.Push(item)

	peekedItem := stackOfPapers.Peek()

	// Check that peekedItem == item
	if item != peekedItem {
		t.Errorf("Item %v and peeked item %v do not match", item, peekedItem)
	}

	// Check that the stack is not empty
	if len(stackOfPapers.elements) != 1 {
		t.Errorf("Stack should not be empty")
	}

	// Check that the top element is still item
	if item != stackOfPapers.elements[0] {
		t.Errorf("The item should not have been removed from the stack")
	}
}

func TestIsEmpty(t *testing.T) {
	stackOfPapers := SliceStack{}

	// Stack should start as empty
	if stackOfPapers.IsEmpty() != true {
		t.Errorf("Stack should be empty, but got %v instead", stackOfPapers)
	}

	stackOfPapers.Push("Business card")

	// Stack not be empty if an item is pushed to it
	if stackOfPapers.IsEmpty() == true {
		t.Errorf("Stack should not be empty, but got %v instead", stackOfPapers)
	}
}

func TestClear(t *testing.T) {
	stackOfPapers := SliceStack{}

	stackOfPapers.Push("Old movie tickets")
	stackOfPapers.Clear()

	// Should have no elements after clearing
	if len(stackOfPapers.elements) != 0 {
		t.Errorf("Stack should have no elements, instead it has %v", stackOfPapers)
	}
}
