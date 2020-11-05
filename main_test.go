package main

import (
	"math/rand"
	"testing"
)

func Test_initialize(test *testing.T) {
	rand.Seed(1)

	result := initialize(5)
	if result != "OPCLE" {
		test.Fail()
	}
}

func Test_mutate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	result := mutate(text, 0.2)

	wantedResult := "the qu Pk brown fox jumps oveF tGD Nazy dog"
	if result != wantedResult {
		test.Fail()
	}
}
