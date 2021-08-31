package weasel

import (
	"math/rand"
)

func mutate(text string, rate float64) string {
	result := ""
	for _, textCharacter := range text {
		if rand.Float64() < rate {
			randomCharacter := makeRandomCharacter()
			result = result + string(randomCharacter)
		} else {
			result = result + string(textCharacter)
		}
	}

	return result
}

func fitness(text string, sample string) int {
	count := 0
	for index, textCharacter := range text {
		if sample[index] != byte(textCharacter) {
			count = count + 1
		}
	}

	return count
}

// Populate ...
func Populate(text string, rate float64, count int) []string {
	textCopies := []string{}
	for i := 0; i < count; i = i + 1 {
		textCopy := mutate(text, rate)
		textCopies = append(textCopies, textCopy)
	}

	return textCopies
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
