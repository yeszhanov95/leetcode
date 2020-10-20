package main

//some base cases
func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) { return arr }
	if arr[0] >= x { return arr[:k] }
	if arr[len(arr)-1] <= x { return arr[len(arr)-k:] }

	//binary search
	l, r, m := 0, len(arr) - 1, 0
	for l < r {
		m = (l + r) / 2
		if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}

	low, high := l, l
	// if there is no match for X in array, decide which index is closest
	if arr[l] != x && (x - arr[l-1]) <= (arr[l] - x) {
		low, high = l - 1, l - 1
	}

	//expansion of array using two pointers
	for ; k > 1; k-- {
		if low == 0 {
			high++
		} else if high == len(arr) - 1 {
			low--
		} else if (arr[high+1] - x) < (x - arr[low-1]) {
			high++
		} else {
			low--
		}
	}

	return arr[low:high+1]
}