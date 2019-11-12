package main

import "fmt"

func lengthOfLastWord(s string) int {
	word := ""
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " && len(word) != 0 {
			break
		}
		if string(s[i]) == " " && len(word) == 0 {
			continue
		}
		word += string(s[i])
	}
	return len(word)
}

func main() {
	fmt.Println(lengthOfLastWord("  asdas  "))
}
