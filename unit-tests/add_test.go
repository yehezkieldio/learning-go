package unittests

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T)  {
	result := Add(1, 2)
	expected := 3

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}