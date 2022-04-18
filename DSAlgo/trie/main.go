package main

import (
	"bufio"
	"fmt"
	"github.com/hiteshrepo/learninggist/dsalgo/trie/trie_data_structure"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	// allPrefixes := getAllPrefixes()
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

func getAllPrefixes() []string {
	allPrefixes := make([]string, 0)
	wd, err := os.Getwd()
	pathToFile := path.Join(wd, "DSAlgo/trie/dict/sample_prefixes.txt")
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allPrefixes = append(allPrefixes, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allPrefixes
}
