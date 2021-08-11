package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func makeRandomCharacter() byte {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}

func initialize(length int) string {
	result := ""
	for i := 0; i < length; i = i + 1 {
		randomCharacter := makeRandomCharacter()
		result = result + string(randomCharacter)
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
	for i := 0; i < count; i = i + 1 {
		textCopy := mutate(text, rate)
		textCopies = append(textCopies, textCopy)
	}

	return textCopies
}

func search(variants []string, sample string) string {
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

func main() {
	rand.Seed(time.Now().UnixNano())

	sample :=
		flag.String("sample", "METHINKS IT IS LIKE A WEASEL", "target string")
	rate := flag.Float64("rate", 0.05, "character mutate rate")
	populationCount := flag.Int("count", 100, "population size")
	flag.Parse()

	generation := 0
	current := initialize(len(*sample))
	for current != *sample {
		const outputRate = 10
		if generation%outputRate == 0 {
			fmt.Println(generation, current)
		}

		variants := populate(current, *rate, *populationCount)
		current = search(variants, *sample)

		generation = generation + 1
	}

	fmt.Println(generation, current)
}
