package distinct_substrings

import (
	"strings"
)

type TrieMapNode struct {
	alphabet map[int]*TrieMapNode
}

func GetNewTrieMapNode() *TrieMapNode {
	return &TrieMapNode{
		alphabet: make(map[int]*TrieMapNode),
	}
}

func (t *TrieMapNode) containsKey(char rune) bool {
	ascii := int(char)
	return t.alphabet[ascii] != nil
}

func (t *TrieMapNode) get(char rune) *TrieMapNode {
	ascii := int(char)
	return t.alphabet[ascii]
}

func (t *TrieMapNode) put(char rune, node *TrieMapNode) {
	ascii := int(char)
	t.alphabet[ascii] = node
}

type MapTrie struct {
	root *TrieMapNode
}

func GetNewMapTrie() *MapTrie {
	return &MapTrie{root: GetNewTrieMapNode()}
}

func (t *MapTrie) Insert(word string) {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _,letter := range letters {
		if !node.containsKey(letter) {
			node.put(letter, GetNewTrieMapNode())
		}

		node = node.get(letter)
	}
}

func (t *MapTrie) PrintDistinctSubstringsOptimal(word string) {
	for i:=0; i<len(word); i++ {
		currentWord := word[i:]
		t.Insert(currentWord)
	}
}
