package main

func countPrimes(n int) int {
	if n < 3 { return 0 }
	slice, nonPrime := make([]byte, n), 2
	for i := 2; i < n; i++ {
		if slice[i] == 0 {
			for j := i * 2; j < n; j += i {
				if slice[j] == 0 { nonPrime++ }
				slice[j] = 1
			}
		}
	}
	return n - nonPrime
}