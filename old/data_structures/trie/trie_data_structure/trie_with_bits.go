package trie_data_structure

type BitsNode struct {
	Links []*BitsNode
}

func GetBitsNode() *BitsNode {
	return &BitsNode{Links: make([]*BitsNode, 2)}
}

func (node *BitsNode) containsKey(bit int) bool {
	return node.Links[bit] != nil
}

func (node *BitsNode) get(bit int) *BitsNode {
	return node.Links[bit]
}

func (node *BitsNode) put(bit int, bn *BitsNode) {
	node.Links[bit] = bn
}

type BitsNodeTrie struct {
	Root *BitsNode
}

func GetBitsNodeTrie() *BitsNodeTrie {
	return &BitsNodeTrie{Root: GetBitsNode()}
}

func (trie *BitsNodeTrie) Insert(num int) {
	node := trie.Root
	for i:=31; i>=0; i-- {
		bit := (num >> i) & 1
		if !node.containsKey(bit) {
			node.put(bit, GetBitsNode())
		}
		node = node.get(bit)
	}
}

func (trie *BitsNodeTrie) GetMax(num int) int {
	node := trie.Root
	maxNum := 0
	for i:=31; i>=0; i-- {
		bit := (num >> i) & 1
		if node.containsKey(1 - bit) {
			maxNum = maxNum | (1 << i)
			node = node.get(1 - bit)
		} else {
			node = node.get(bit)
		}
	}
	return maxNum
}
