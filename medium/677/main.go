package main

type MapSum struct {
	root *TrieNode
}

type TrieNode struct {
	chars [26]*TrieNode
	val int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	root := &TrieNode{}
	return MapSum{root}
}

func (this *MapSum) Insert(key string, val int)  {
	curr := this.root
	for i := range key {
		char := key[i] - 'a'
		if curr.chars[char] == nil {
			curr.chars[char] = &TrieNode{}
		}
		curr = curr.chars[char]
	}
	curr.val = val
}

func (this *MapSum) Sum(prefix string) int {
	curr := this.root
	for i := range prefix {
		char := prefix[i] - 'a'
		if curr.chars[char] == nil {
			return 0
		}
		curr = curr.chars[char]
	}
	result := 0
	sum(curr, &result)
	return result
}

func sum(node *TrieNode, total *int) {
	*total += node.val
	for _, v := range node.chars {
		if v != nil {
			sum(v, total)
		}
	}
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */