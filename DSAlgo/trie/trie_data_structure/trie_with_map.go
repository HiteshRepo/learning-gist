package trie_data_structure

import "strings"

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

	for _,letter := range letters {
		if !node.containsKey(letter) {
			node.put(letter, GetNewTrieMapNode())
		}

		node = node.get(letter)
	}

	node.setEnd()
}


func (t *MapTrie) Search(word string) bool {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _,letter := range letters {
		if !node.containsKey(letter) {
			return false
		}
		node = node.get(letter)
	}

	return node.isEnd()
}

func (t *MapTrie) StartsWith(prefix string) bool {
	node := t.root
	copyWord := strings.ToLower(prefix)
	letters := []rune(copyWord)

	for _,letter := range letters {
		if !node.containsKey(letter) {
			return false
		}
		node = node.get(letter)
	}

	return true
}

func (t *MapTrie) GetLongestWord(words []string) string {
	completeString := ""

	for _,w := range words {
		if !t.Search(w) {
			continue
		}

		if len(w) > len(completeString) {
			completeString = w
		}
	}

	return completeString
}
