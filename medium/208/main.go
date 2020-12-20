package main

type Trie struct {
	root *TrieNode
}


type TrieNode struct {
	chars [26]*TrieNode
	isWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	root := &TrieNode{chars: [26]*TrieNode{}}
	return Trie{root}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	curr := this.root
	for i := range word {
		char := word[i]-'a'
		if curr.chars[char] == nil {
			curr.chars[char] = &TrieNode{chars: [26]*TrieNode{}}
		}
		curr = curr.chars[char]
	}
	curr.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := range word {
		char := word[i]-'a'
		if curr.chars[char] == nil {
			return false
		}
		curr = curr.chars[char]
	}
	return curr.isWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := range prefix {
		char := prefix[i]-'a'
		if curr.chars[char] == nil {
			return false
		}
		curr = curr.chars[char]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */