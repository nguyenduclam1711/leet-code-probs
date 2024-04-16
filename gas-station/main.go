package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank, tankLeft, start := 0, 0, 0

	for i := range gas {
		totalTank += gas[i] - cost[i]
		tankLeft += gas[i] - cost[i]

		if tankLeft < 0 {
			start = i + 1
			tankLeft = 0
		}
	}
	if totalTank < 0 {
		return -1
	}
	return start
}
