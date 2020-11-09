package main

import "math/rand"

func makeRandomCharacter() byte {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}

func initialize(length int) string {
	result := ""
	i := 0
	for i < length {
		randomCharacter := makeRandomCharacter()
		result = result + string(randomCharacter)
		i = i + 1
	}

	return result
}

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

func populate(text string, rate float64, count int) []string {
	textCopies := []string{}
	i := 0
	for i < count {
		textCopy := mutate(text, rate)
		textCopies = append(textCopies, textCopy)
		i = i + 1
	}

	return textCopies
}
