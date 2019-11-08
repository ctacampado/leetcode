package main

import (
	"fmt"
	"math"
)

//Reverse Given a 32-bit signed integer, reverse digits of an integer.
/*
Example 1:
Input: 123
Output: 321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
func Reverse(x int) int {

	s := 1
	var res, in int

	if x > math.MaxInt32 || x < math.MinInt32 || x == 0 {
		return 0
	}

	if x < 0 {
		s *= -1
	}

	in = x * s
	length := int(math.Log10(float64(in))) + 1

	for i := 0; i < length; i++ {
		res = res*10 + in%10
		in /= 10
		if res > math.MaxInt32 {
			return 0
		}
	}

	return res * s
}

func main() {
	fmt.Println(Reverse(-2))
}
