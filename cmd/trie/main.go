package main

import (
	"fmt"
	"strings"
)

const AlphabethSize = 26

type TrieNode struct {
	children [AlphabethSize]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

// Create a new Trie
func InitializeTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

// Add word to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := range wordLength {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &TrieNode{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search returns true if the word is found in the trie
func (t *Trie) Search(w string) bool {
	lcWord := strings.ToLower(w)
	wordLength := len(lcWord)
	currentNode := t.root
	for i := range wordLength {
		charIndex := lcWord[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func main() {
	testTrie := InitializeTrie()

	testSlice := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
	}

	for _, v := range testSlice {
		testTrie.Insert(v)
	}

	fmt.Println(testTrie.Search("Aragorn"))
}
