package main

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	trie := initTrie(dictionary)
	words := strings.Fields(sentence)
	result := getRoot(&trie, words[0])
	for i := 1; i < len(words); i++ {
		result += " " + getRoot(&trie, words[i])
	}
	return result
}

func initTrie(words []string) Trie {
	trie := Trie{root: &TrieNode{}}
	for _, word := range words {
		curr := trie.root
		for i := range word {
			char := word[i]-'a'
			if curr.children[char] == nil {
				curr.children[char] = &TrieNode{}
			}
			curr = curr.children[char]
			if curr.isWord {
				break
			}
		}
		curr.isWord = true
	}
	return trie
}

func getRoot(trie *Trie, word string) string {
	curr := trie.root
	for i := range word {
		char := word[i]-'a'
		if curr.children[char] == nil { return word }
		curr = curr.children[char]
		if curr.isWord { return word[:i+1] }
	}
	return word
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	isWord bool
}
