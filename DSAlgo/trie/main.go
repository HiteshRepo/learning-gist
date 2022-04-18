package main

import (
	"fmt"
	"github.com/hiteshrepo/learninggist/dsalgo/trie/trie_data_structure"
	"log"
)

func main() {
	log.Println("--------Trie---------")
	TestTrie()
	log.Println("--------Trie With Map---------")
	TestTrieWithMap()
}

func TestTrieWithMap() {
	allPrefixes := getTestPrefixes()

	trie := trie_data_structure.GetNewMapTrie()

	for _, word := range allPrefixes {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("apps"))
	fmt.Println(trie.Search("bac"))
	fmt.Println(trie.Search("azer"))
	fmt.Println(trie.StartsWith("ap"))
	fmt.Println(trie.StartsWith("ba"))
	fmt.Println(trie.StartsWith("l"))
}

func TestTrie() {
	allPrefixes := getTestPrefixes()

	trie := trie_data_structure.GetNewTrie()

	for _, word := range allPrefixes {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("apps"))
	fmt.Println(trie.Search("bac"))
	fmt.Println(trie.Search("azer"))
	fmt.Println(trie.StartsWith("ap"))
	fmt.Println(trie.StartsWith("ba"))
	fmt.Println(trie.StartsWith("l"))
}

func getTestPrefixes() []string {
	return []string{"apple", "apps", "apxl", "bac", "bat"}
}
