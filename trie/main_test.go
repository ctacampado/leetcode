package main

import "testing"

type TestCase struct {
	input []string
	res   string
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []TestCase{
		{
			input: []string{"flower", "flow", "floight"},
			res:   "flo",
		}, {
			input: []string{"dog", "racecar", "car"},
			res:   "",
		}, {
			input: []string{"a"},
			res:   "a",
		}, {
			input: []string{""},
			res:   "",
		}, {
			input: []string{"aa", "aa"},
			res:   "aa",
		}, {
			input: []string{"a", "b"},
			res:   "",
		}, {
			input: []string{"c", "c"},
			res:   "c",
		}, {
			input: []string{"c", "c", "c"},
			res:   "c",
		}, {
			input: []string{"c", "c", "b"},
			res:   "",
		}, {
			input: []string{"aaa", "aa", "aaa"},
			res:   "aa",
		}, {
			input: []string{"aaa", "aaa", "aa"},
			res:   "aa",
		}, {
			input: []string{"aa", "aaa", "aaa"},
			res:   "aa",
		}, {
			input: []string{"", "aba", "abab"},
			res:   "",
		}, {
			input: []string{"c", "acc", "ccc"},
			res:   "",
		}, {
			input: []string{"caa", "", "a", "acb"},
			res:   "",
		}, {
			input: []string{"aac", "acab", "aa", "abba", "aa"},
			res:   "a",
		}, {
			input: []string{"aa", "ab"},
			res:   "a",
		}, {
			input: []string{"abab", "aba", ""},
			res:   "",
		},
	}

	for _, tc := range testCases {
		res := LongestCommonPrefix(tc.input)
		if res != tc.res {
			t.Errorf("incorrect result. exp %s | got %s\n", tc.res, res)
		} else {
			t.Logf("passed. exp %s | got %s\n", tc.res, res)
		}
	}
}
