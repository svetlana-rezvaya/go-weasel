package weasel

import "math/rand"

func makeRandomCharacter() byte {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}

// Initialize ...
func Initialize(length int) string {
	result := ""
	for i := 0; i < length; i = i + 1 {
		randomCharacter := makeRandomCharacter()
		result = result + string(randomCharacter)
	}

	return result
}
