package main

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	isWord bool
}

func Constructor() WordDictionary {
	root := &TrieNode{}
	return WordDictionary{root}
}

func (this *WordDictionary) AddWord(word string)  {
	curr := this.root
	for i := range word {
		char := word[i] - 'a'
		if curr.children[char] == nil {
			curr.children[char] = &TrieNode{}
		}
		curr = curr.children[char]
	}
	curr.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.root, []byte(word))
}

func dfs(node *TrieNode, chars []byte) bool {
	for i := range chars {
		if chars[i] == '.' {
			for _, child := range node.children {
				if child != nil && dfs(child, chars[i+1:]) {
					return true
				}
			}
			return false
		} else {
			c := chars[i] - 'a'
			if node.children[c] == nil {
				return false
			}
			node = node.children[c]
		}
	}
	return node.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */