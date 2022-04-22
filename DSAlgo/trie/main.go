package main

import (
	"fmt"
	"github.com/hiteshrepo/learninggist/dsalgo/trie/distinct_substrings"
	"github.com/hiteshrepo/learninggist/dsalgo/trie/trie_data_structure"
	"log"
)

func main() {
	log.Println("--------Trie---------")
	TestTrie()
	log.Println("--------Trie With Map---------")
	TestTrieWithMap()
	log.Println("--------Trie Ce Cp---------")
	TestCeCpTrie()
	log.Println("--------Longest Word---------")
	TestLongestWord()
	log.Println("--------Distinct Substrings---------")
	TestDistinctSubstrings()
}

func TestDistinctSubstrings() {
	distinct_substrings.PrintDistinctSubstringsCount("abab")
}

func TestLongestWord() {
	allWords := []string{"n", "ninja", "ni", "nin", "ninjaa", "ninga"}

	trie := trie_data_structure.GetNewMapTrie()

	for _, word := range allWords {
		trie.Insert(word)
	}

	findWords := []string{"n", "ni", "nin", "ning"}

	fmt.Println(trie.GetLongestWord(findWords)) //nin
}

func TestCeCpTrie() {
	allPrefixes := getTestPrefixes()

	trie := trie_data_structure.GetNewCeCpTrie()

	for _, word := range allPrefixes {
		trie.Insert(word)
	}

	fmt.Println(trie.CountWordsStartsWith("ap")) //3
	fmt.Println(trie.CountWordsStartsWith("ba")) //2
	fmt.Println(trie.CountWordsStartsWith("az")) //0
	fmt.Println(trie.CountWordsEndWith("apple")) //1
	fmt.Println(trie.CountWordsEndWith("bac")) //1
	fmt.Println(trie.CountWordsEndWith("vat")) //0

	trie.Erase("bac")
	fmt.Println(trie.CountWordsEndWith("bac")) //0
	trie.Erase("apple")
	fmt.Println(trie.CountWordsStartsWith("ap")) //2
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
