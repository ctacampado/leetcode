/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true

Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Follow up:

Coud you solve it without converting the integer to a string?
*/
package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x <= 0 {
		return x == 0
	}
	length := int(math.Log10(float64(x)))
	mask := int(math.Pow10(length))
	for i := 0; i <= length/2; i++ {
		if x%10 != x/mask {
			return false
		}
		x %= mask
		x /= 10
		mask /= 100
	}
	return true
}

func isPalindrome2(x int) bool {

	length := int(math.Log10(float64(x)))
	ln, rn := 0, 0
	y1, y2 := x, x

	for i := 0; i <= length/2; i++ {
		ln = y1 - (y1 - int(math.Pow10(length)))
		rn = y2 % 10 * int(math.Pow10(length))
		if ln != rn {
			return false
		}
		y1 = y1 / ln
		y2 = y2 % 10
	}
	return true
}

func main() {
	//print(isPalindrome(1000110001))
	//fmt.Println(isPalindrome1(115111))
	fmt.Println(isPalindrome2(1122131))
}
