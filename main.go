package weasel

func fitness(text string, sample string) int {
	count := 0
	for index, textCharacter := range text {
		if sample[index] != byte(textCharacter) {
			count = count + 1
		}
	}

	return count
}

// Search ...
func Search(variants []string, sample string) string {
	result := ""
	minCount := len(sample)
	for _, variant := range variants {
		count := fitness(variant, sample)
		if count < minCount {
			result = variant
			minCount = count
		}
	}

	return result
}
