package main

func canCompleteCircuit(gas []int, cost []int) int {
	var diff int
	for i := 0; i < len(gas); i++ {
		diff += gas[i] - cost[i]
	}
	if diff < 0 { return -1 }

	var start, tank int
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	return start
}