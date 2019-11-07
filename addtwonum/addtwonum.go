/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package main

import (
	"fmt"
)

//ListNode node struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ref := ListNode{}
	head := &ref
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{Val: sum % 10}
		head = head.Next
		sum /= 10
	}
	if sum != 0 {
		head.Next = &ListNode{Val: sum}
	}
	return ref.Next
}

func main() {
	//add 123 456
	l1c := ListNode{Val: 3, Next: nil}
	l1b := ListNode{Val: 4, Next: &l1c}
	l1 := ListNode{Val: 2, Next: &l1b}

	l2c := ListNode{Val: 4, Next: nil}
	l2b := ListNode{Val: 6, Next: &l2c}
	l2 := ListNode{Val: 5, Next: &l2b}

	res := addTwoNumbers(&l1, &l2)

	currNodel1 := res
	n1 := []int{}
	for {
		n1 = append(n1, currNodel1.Val)
		if currNodel1.Next != nil {
			currNodel1 = currNodel1.Next
		} else {
			break
		}
	}
	fmt.Println(n1)
}
