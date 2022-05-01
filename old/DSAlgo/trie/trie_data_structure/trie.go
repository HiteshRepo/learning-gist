package trie_data_structure

import (
	"strings"
)

const (
	StartAscii = 97
)

type TrieNode struct {
	alphabet []*TrieNode
	last     bool
}

func GetNewTrieNode() *TrieNode {
	return &TrieNode{
		alphabet: make([]*TrieNode, 26),
		last:     false,
	}
}

func (t *TrieNode) containsKey(char rune) bool {
	ascii := int(char)
	position := ascii - StartAscii
	return t.alphabet[position] != nil
}

func (t *TrieNode) get(char rune) *TrieNode {
	ascii := int(char)
	position := ascii - StartAscii
	return t.alphabet[position]
}

func (t *TrieNode) put(char rune, node *TrieNode) {
	ascii := int(char)
	position := ascii - StartAscii
	t.alphabet[position] = node
}

func (t *TrieNode) setEnd() {
	t.last = true
}

func (t *TrieNode) isEnd() bool {
	return t.last
}

type Trie struct {
	root *TrieNode
}

func GetNewTrie() *Trie {
	return &Trie{root: GetNewTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _,letter := range letters {
		if !node.containsKey(letter) {
			node.put(letter, GetNewTrieNode())
		}

		node = node.get(letter)
	}

	node.setEnd()
}


func (t *Trie) Search(word string) bool {
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

func (t *Trie) StartsWith(prefix string) bool {
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