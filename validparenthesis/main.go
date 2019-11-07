package main

import "fmt"

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true

Example 2:

Input: "()[]{}"
Output: true

Example 3:

Input: "(]"
Output: false

Example 4:

Input: "([)]"
Output: false

Example 5:

Input: "{[]}"
Output: true
*/

//Stack data struct for a FIFO stack
type Stack []string
type LookUp map[string]int

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() string {
	l := len(*s) - 1
	elem := (*s)[l]
	*s = (*s)[:l]
	return elem
}

func (s *Stack) top() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func checkIfOpen(s string) bool {
	return s == "{" || s == "(" || s == "["
}

func checkIfMatch(open, close string) bool {
	if open == "{" && close == "}" ||
		open == "[" && close == "]" ||
		open == "(" && close == ")" {
		return true
	}
	return false
}
func isValidA(s string) bool {
	var pStack Stack
	for _, e := range s {
		v := string(e)
		if !checkIfOpen(v) && !pStack.isEmpty() {
			if checkIfMatch(pStack.top(), v) {
				pStack.pop()
			} else {
				return false
			}
		} else {
			pStack.push(v)
		}
	}
	return pStack.isEmpty()
}

func isValid(s string) bool {
	oLookUp := LookUp{
		"{": 1,
		"(": 2,
		"[": 3,
	}
	cLookUp := LookUp{
		"}": 1,
		")": 2,
		"]": 3,
	}
	var pStack Stack
	for _, e := range s {
		v := string(e)
		if _, ok := oLookUp[v]; !ok && !pStack.isEmpty() {
			if oLookUp[pStack.top()] == cLookUp[v] {
				pStack.pop()
			} else {
				return false
			}
		} else {
			pStack.push(v)
		}
	}
	return pStack.isEmpty()
}

func main() {
	fmt.Println(isValid("(([]){})"))
}
