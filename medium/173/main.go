package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
type BSTIterator struct {
	data []int
	currIndex int
}


func Constructor(root *TreeNode) BSTIterator {
	stack, data := make([]*TreeNode, 0), make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		data = append(data, root.Val)
		root = root.Right
	}
	return BSTIterator{data: data, currIndex: -1}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.currIndex++
	return this.data[this.currIndex]
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.currIndex < (len(this.data) - 1)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */