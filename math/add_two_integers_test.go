package leetcode

import "testing"

func TestAddTwoIntegers(t *testing.T) {
	output := sum(2, 2)
	expected := 4

	if output != expected {
		t.Errorf("got %d expeceted %d", output, expected)
	}
}
