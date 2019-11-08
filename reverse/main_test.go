package main

import (
	"math"
	"testing"
)

type TestCase struct {
	input int
	res   int
}

func TestReverse(t *testing.T) {
	testCases := []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
		{math.MaxInt32 + 1, 0},
		{math.MinInt32 - 1, 0},
	}

	for _, tc := range testCases {
		res := Reverse(tc.input)
		if res != tc.res {
			t.Errorf("incorrect result. exp %d | got %d\n", tc.res, res)
		} else {
			t.Logf("passed. exp %d | got %d\n", tc.input, res)
		}
	}
}
