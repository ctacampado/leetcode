package searchsort

import (
	"reflect"
	"testing"
)

func TestBSearch(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	target := 4
	answer := 3
	i := BSearch(input, target)
	if i != answer {
		t.Errorf("answer was incorrect, got: %d, want: %d.", i, answer)
	}
}

func TestSortTwoSorted(t *testing.T) {
	input1 := []int{2, 4, 16}
	input2 := []int{1, 3, 7, 9, 21}
	answer := []int{1, 2, 3, 4, 7, 9, 16, 21}
	res := SortTwoSorted(input1, input2)
	if !reflect.DeepEqual(res, answer) {
		t.Errorf("answer was incorrect, got: %+v, want: %+v.", res, answer)
	}
}
