/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import (
	"fmt"
)

//TwoSum -
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
func TwoSum(nums []int, target int) []int {
	nIdx := len(nums) - 1
	sMap := map[int]int{}
	for i, v := range nums {
		j, ok := sMap[target-v]
		if ok && i != j {
			return []int{i, j}
		}
		sMap[v] = i
		k, ok := sMap[target-nums[nIdx]]
		if ok && ((nIdx) != k) {
			return []int{nIdx, k}
		}
		sMap[nums[nIdx]] = nIdx
		nIdx--
	}
	return nil
}

func main() {
	test1Nums := []int{2, -1, 11, 15, 1, -1, -1, 11, 15, 1, -1, -1, 11, 15, 1, -1, -2}
	test1Targ := 9
	res := TwoSum(test1Nums, test1Targ)
	fmt.Println(res)

}
