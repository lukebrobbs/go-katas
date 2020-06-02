package splitandadd

import (
	"reflect"
	"testing"
)

type Test struct {
	run      string
	expected []int
	actual   []int
}

func TestSplitAndAdd(t *testing.T) {
	tests := []Test{
		{
			run:      "if n is 1, should only do the process once",
			expected: []int{3, 6, 7, 12},
			actual:   splitAndAdd([]int{4, 2, 5, 3, 2, 5, 7}, 1),
		},
		{
			run:      "Should handle n being more than 1",
			expected: []int{10, 18},
			actual:   splitAndAdd([]int{4, 2, 5, 3, 2, 5, 7}, 2),
		},
		{
			run:      "numbers single length slice",
			expected: []int{1},
			actual:   splitAndAdd([]int{1}, 5),
		},
	}

	for _, te := range tests {
		t.Run(te.run, func(t *testing.T) {
			if !reflect.DeepEqual(te.actual, te.expected) {
				t.Errorf("Expected %v, instead got %v", te.expected, te.actual)
			}
		})
	}
}
