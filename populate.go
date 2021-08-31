package weasel

import "math/rand"

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

// Populate ...
func Populate(text string, rate float64, count int) []string {
	textCopies := []string{}
	for i := 0; i < count; i = i + 1 {
		textCopy := mutate(text, rate)
		textCopies = append(textCopies, textCopy)
	}

	return textCopies
}
