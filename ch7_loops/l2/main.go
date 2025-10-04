package main

func maxMessages(thresh int) int {
	totalCost := 0
	maxVal := 0
	for i := 0; ; i++ {
		totalCost += (100 + i)
		if totalCost > thresh {
			maxVal = i
			break
		}
	}
	return maxVal
}
