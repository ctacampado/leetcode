package main

import (
	"testing"
)

type Input struct {
	arr []int
	t   int
}

type TestCase struct {
	input Input
	res   []int
}

func isEqualUnordered(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	//// create a map of int -> int
	diff := make(map[int]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}
	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{
			input: Input{
				arr: []int{2, 7, 11, 15},
				t:   9,
			},
			res: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		res := TwoSum(tc.input.arr, tc.input.t)
		if !isEqualUnordered(res, tc.res) {
			t.Errorf("incorrect result. exp %v | got %v\n", tc.res, res)
		} else {
			t.Logf("passed. exp %v | got %v\n", tc.res, res)
		}
	}
}
