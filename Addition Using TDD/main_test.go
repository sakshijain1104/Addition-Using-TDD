package main

import "testing"

type AddResult struct {
	x              int
	y              int
	expectedOutput int
}

var AddResults = []AddResult{
	{2, 3, 5},
	{1, 1, 2},
}

func TestAdd(t *testing.T) {
	for _, test := range AddResults {
		result := Add(test.x, test.y)
		if result != test.expectedOutput {
			t.Error("not the expected reult")
		}
	}
}
