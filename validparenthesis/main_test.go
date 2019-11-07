package main

import "testing"

type TestCase struct {
	input string
	res   bool
}

func TestIsValid(t *testing.T) {
	testCases := []TestCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"([)", false},
		{"(])", false},
		{"(([]){})", true},
	}

	for _, tc := range testCases {
		res := isValid(tc.input)
		if res != tc.res {
			t.Errorf("incorrect result. expected %v | got %v\n", tc.res, res)
		} else {
			t.Logf("passed. expected %v | got %v", tc.res, res)
		}
	}
}
