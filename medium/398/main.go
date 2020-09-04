package main

type Solution struct {
	cache map[int][]int
	lastIdx int
}


func Constructor(nums []int) Solution {
	cache := make(map[int][]int, len(nums))
	for i, v := range nums {
		if _, ok := cache[v]; ok {
			cache[v] = append(cache[v], i)
		} else {
			cache[v] = []int{i}
		}
	}
	return Solution{cache, 0}
}


func (this *Solution) Pick(target int) int {
	slice := this.cache[target]
	if this.lastIdx == len(slice) - 1 {
		this.lastIdx = 0
	} else {
		this.lastIdx++
	}
	return slice[this.lastIdx]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */