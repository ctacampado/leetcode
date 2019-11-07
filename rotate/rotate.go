package main

import (
	"fmt"
	"strconv"
)

func rotate(N int) int {
	numStr := strconv.Itoa(N)
	tStr := []byte{}
	for _, v := range numStr {
		if v == '2' || v == '5' || v == '6' || v == '9' || v == '8' || v == '0' || v == '1' {
			e := v
			if v == '2' {
				e = '5'
			}
			if v == '5' {
				e = '2'
			}
			if v == '6' {
				e = '9'
			}
			if v == '9' {
				e = '6'
			}
			tStr = append(tStr, byte(e))
		} else {
			return -1
		}
	}
	r, _ := strconv.Atoi(string(tStr[:]))
	return r
}

func eval(N int) bool {
	r := rotate(N)
	fmt.Println("R: ", r)
	fmt.Println("N: ", N)
	if N != r && r != -1 {
		return true
	}
	return false
}

func rotatedDigits(N int) int {
	ctr := 0
	for i := 1; i <= N; i++ {
		if eval(i) {
			ctr++
		}
	}
	return ctr
}

func main() {

	fmt.Print("\n", rotatedDigits(20))
}
