package main

import "math/rand"

func initialize(length int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	result := ""
	i := 0
	for i < length {
		randomIndex := rand.Intn(len(alphabet))
		character := alphabet[randomIndex]
		result = result + string(character)
		i = i + 1
	}

	return result
}

func mutate(text string, rate float64) string {
	result := ""
	for _, character := range text {
		result = result + string(character)
	}

	return result
}
