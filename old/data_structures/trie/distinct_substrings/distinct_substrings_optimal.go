package distinct_substrings

import (
	"fmt"
	"strings"
)

type TrieMapNode struct {
	alphabet map[int]*TrieMapNode
	last     bool
}

func GetNewTrieMapNode() *TrieMapNode {
	return &TrieMapNode{
		alphabet: make(map[int]*TrieMapNode),
		last:     false,
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

func (t *TrieMapNode) setEnd() {
	t.last = true
}

func (t *TrieMapNode) isEnd() bool {
	return t.last
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

	for _, letter := range letters {
		if !node.containsKey(letter) {
			node.put(letter, GetNewTrieMapNode())
		}

		node = node.get(letter)
	}

	node.setEnd()
}

func (t *MapTrie) PrintDistinctSubstringsOptimal(word string) {
	for i := 0; i < len(word); i++ {
		currentWord := word[i:]
		t.Insert(currentWord)
	}

	substrs := make([]string, 0)
	t.Display("", t.root, &substrs)
	fmt.Println(substrs)
}

func (t *MapTrie) Display(prefix string, node *TrieMapNode, substrs *[]string) {
	if node.isEnd() {
		*substrs = append(*substrs, prefix)
	}

	for k,n := range node.alphabet {
		newPrefix := fmt.Sprintf("%s%s", prefix, string(rune(k)))
		t.Display(newPrefix, n, substrs)
	}
}
