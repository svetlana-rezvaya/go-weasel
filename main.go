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
