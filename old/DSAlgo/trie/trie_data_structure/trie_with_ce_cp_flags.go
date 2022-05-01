package trie_data_structure

import "strings"

type TrieCeCpNode struct {
	alphabet    map[int]*TrieCeCpNode
	countEnd    int
	countPrefix int
}

func GetNewTrieCeCpNode() *TrieCeCpNode {
	return &TrieCeCpNode{
		alphabet:    make(map[int]*TrieCeCpNode),
		countEnd:    0,
		countPrefix: 0,
	}
}

func (t *TrieCeCpNode) containsKey(char rune) bool {
	ascii := int(char)
	return t.alphabet[ascii] != nil
}

func (t *TrieCeCpNode) get(char rune) *TrieCeCpNode {
	ascii := int(char)
	return t.alphabet[ascii]
}

func (t *TrieCeCpNode) put(char rune, node *TrieCeCpNode) {
	ascii := int(char)
	t.alphabet[ascii] = node
}

func (t *TrieCeCpNode) incEnd() {
	t.countEnd += 1
}

func (t *TrieCeCpNode) decEnd() {
	t.countEnd -= 1
}

func (t *TrieCeCpNode) isEnd() bool {
	return t.countEnd > 0
}

func (t *TrieCeCpNode) incPrefix() {
	t.countPrefix += 1
}

func (t *TrieCeCpNode) decPrefix() {
	t.countPrefix -= 1
}

func (t *TrieCeCpNode) getPrefix() int {
	return t.countPrefix
}

func (t *TrieCeCpNode) getEnd() int {
	return t.countEnd
}

type CeCpTrie struct {
	root *TrieCeCpNode
}

func GetNewCeCpTrie() *CeCpTrie {
	return &CeCpTrie{root: GetNewTrieCeCpNode()}
}

func (t *CeCpTrie) Insert(word string) {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _, letter := range letters {
		if !node.containsKey(letter) {
			n := GetNewTrieCeCpNode()
			n.incPrefix()
			node.put(letter, n)
		} else {
			nextNode := node.get(letter)
			nextNode.incPrefix()
		}

		node = node.get(letter)
	}

	node.incEnd()
}

func (t *CeCpTrie) CountWordsEndWith(word string) int {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _, letter := range letters {
		if !node.containsKey(letter) {
			return 0
		}

		node = node.get(letter)
	}

	return node.getEnd()
}

func (t *CeCpTrie) CountWordsStartsWith(word string) int {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _, letter := range letters {
		if !node.containsKey(letter) {
			return 0
		}

		node = node.get(letter)
	}

	return node.getPrefix()
}

func (t *CeCpTrie) Erase(word string) {
	node := t.root
	copyWord := strings.ToLower(word)
	letters := []rune(copyWord)

	for _, letter := range letters {
		if node.containsKey(letter) {
			nextNode := node.get(letter)
			nextNode.decPrefix()
		}

		node = node.get(letter)
	}

	node.decEnd()
}
