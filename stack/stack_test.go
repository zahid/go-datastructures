package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stackOfPapers := Stack{}

	stackOfPapers.push("Phone number")

	expectedSize := 1
	actualSize := len(stackOfPapers.elements)

	if expectedSize != actualSize {
		t.Logf("Expected size of %v, but got %v\n", expectedSize, actualSize)
		t.Fail()
	}
}
