package main

// A Trie is comprised of connected Nodes, one Node per char
type Node struct {
	val      interface{}
	children []*Node
}

type Trie struct {
	root Node
}

/*
Question:
Given an input string and a dictionary of words,
find out if the input string can be segmented into a space-separated sequence of dictionary words.

Consider the following dictionary
{ i, like, sam, sung, samsung, mobile, ice,
  cream, icecream, man, go, mango}

	Input:  ilikesamsung
	Output: Yes
	The string can be segmented as
	"i like samsung"
	or "i like sam sung".
*/

func BuildDictionary(words []string) Trie {
	d := Trie{root: Node{val: ""}}
	for _, w := range words {
		for i := 0; i < len(w); i++ {

		}
	}
}
