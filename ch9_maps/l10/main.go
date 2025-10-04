package main

func getNameCounts(names []string) map[rune]map[string]int {
	nestedMap := make(map[rune]map[string]int)

	for _, name := range names {
		for _, ch := range name {
			if _, ok := nestedMap[ch]; !ok {
				nestedMap[ch] = make(map[string]int)
			}

			nestedMap[ch][name] += 1
		}
	}
	return nestedMap
}
