package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	weasel "github.com/svetlana-rezvaya/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sample :=
		flag.String("sample", "METHINKS IT IS LIKE A WEASEL", "target string")
	rate := flag.Float64("rate", 0.05, "character mutate rate")
	populationCount := flag.Int("count", 100, "population size")
	flag.Parse()

	generation := 0
	current := weasel.Initialize(len(*sample))
	for current != *sample {
		const outputRate = 10
		if generation%outputRate == 0 {
			fmt.Println(generation, current)
		}

		variants := weasel.Populate(current, *rate, *populationCount)
		current = weasel.Search(variants, *sample)

		generation = generation + 1
	}

	fmt.Println(generation, current)
}
