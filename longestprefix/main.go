package main

import "fmt"

/*
If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Note:

All given inputs are in lowercase letters a-z.
*/

func maxStr(a, b string) string {
	if len(a) < len(b) {
		return b
	}
	return a
}

func minStr(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//LongestCommonPrefix Write a function to find the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	var minl int
	for i := 1; i < len(strs); i++ {
		minl = minInt(len(prefix), len(strs[i]))
		if minl == 0 {
			return ""
		}
		for j := 0; j < minl; j++ {
			if (prefix)[j] != strs[i][j] {
				prefix = strs[i][:j]
				break
			}
			prefix = minStr(prefix, prefix[:minl])
		}
	}
	return prefix
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"abab", "aba", ""}))
}
