package main

import "fmt"

//TrieNode -
type TrieNode struct {
	tMap map[string]*TrieNode
	end  bool
}

//Insert -
func (tn *TrieNode) Insert(s string) {
	end := false
	head := tn
	for i, v := range s {
		if _, ok := head.tMap[string(v)]; !ok {
			head.tMap[string(v)] = &TrieNode{
				tMap: make(map[string]*TrieNode),
				end:  end,
			}
		}
		if i == len(s)-1 {
			head.tMap[string(v)].end = true
		}
		head = head.tMap[string(v)]
	}
}

//SearchLCP -
func (tn *TrieNode) SearchLCP() string {
	lpc := ""
	head := tn
	for {
		tlen := len(head.tMap)
		//this makes sure that loop after this if-block runs for 1 iteration only
		if head.end || tlen > 1 || tlen < 1 {
			return lpc
		}
		for k := range head.tMap {
			lpc += k
			head = head.tMap[k]
		}
	}
}

//LongestCommonPrefix Write a function to find the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(strs []string) string {

	trie := TrieNode{
		tMap: make(map[string]*TrieNode),
		end:  false,
	}
	if len(strs) == 0 {
		return ""
	}
	for _, v := range strs {
		if len(v) == 0 {
			return ""
		}
		trie.Insert(v)
	}
	return trie.SearchLCP()
}

func main() {
	fmt.Println("lpc: ", LongestCommonPrefix([]string{"flower", "f", "floight"}))
}
