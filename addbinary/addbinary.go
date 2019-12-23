package main

import "strconv"

// AddBinary Given two binary strings, return their sum (also a binary string).
//
// The input strings are both non-empty and contains only characters 1 or 0.
//
// Example 1:
//
// Input: a = "11", b = "1"
// Output: "100"
//
// Example 2:
//
// Input: a = "1010", b = "1011"
// Output: "10101"
func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	carry := 0
	s := ""

	zeroes := ""

	for i := 0; i < len(a)-len(b); i++ {
		zeroes += "0"
	}

	b = zeroes + b

	for i := len(a) - 1; i >= 0; i-- {
		sum := carry
		if string(a[i]) == "1" {
			sum++
		}
		if string(b[i]) == "1" {
			sum++
		}
		carry = sum / 2
		sum = sum % 2
		s = strconv.Itoa(sum) + s
	}

	if carry > 0 {
		s = "1" + s
	}

	return s
}
