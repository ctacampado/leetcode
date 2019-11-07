package main

import (
	"testing"
)

type TestCase struct {
	input string
	res   int
}

func TestLengthOfLongestSubstring(t *testing.T) {

	testCases := []TestCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{" ", 1},
		{"aab", 2},
		{"dvdf", 3},
		{"abba", 2},
	}

	var res int
	for _, tc := range testCases {
		res = LengthOfLongestSubstring(tc.input)
		if res != tc.res {
			t.Errorf("incorrect result. expected %d | got %d\n", tc.res, res)
		} else {
			t.Logf("pass. expected %d | got %d\n", tc.res, res)
		}
	}
}
