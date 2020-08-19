package main

type NumArray struct {
	Nums []int
}


func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{Nums: nums}
}


func (this *NumArray) SumRange(i int, j int) int {
	sum := this.Nums[j]
	if(i > 0) {
		sum -= this.Nums[i-1]
	}
	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */