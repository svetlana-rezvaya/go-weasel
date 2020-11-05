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
