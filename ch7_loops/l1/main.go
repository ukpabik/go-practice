package main

func bulkSend(numMessages int) float64 {
	fee := 0.01
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1 + (fee * float64(i))
	}
	return totalCost
}
