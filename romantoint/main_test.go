package main

import "testing"

type TestCase struct {
	input string
	res   int
}

func TestRomanToInt(t *testing.T) {
	testCases := []TestCase{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		res := RomanToInt(tc.input)
		if res != tc.res {
			t.Errorf("incorrect result. exp %d | got %d\n", tc.res, res)
		} else {
			t.Logf("passed. exp %d | got %d\n", tc.res, res)
		}
	}
}
