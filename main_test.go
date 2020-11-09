package main

import (
	"math/rand"
	"reflect"
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

func Test_fitness_withNotEqualStrings(test *testing.T) {
	count := fitness("test1 test2 test3", "test1 test4 test5")
	if count != 2 {
		test.Fail()
	}
}

func Test_fitness_withEqualStrings(test *testing.T) {
	count := fitness("test1 test2 test3", "test1 test2 test3")
	if count != 0 {
		test.Fail()
	}
}

func Test_populate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	textCopies := populate(text, 0.2, 3)

	wantedTextCopies := []string{
		"the qu Pk brown fox jumps oveF tGD Nazy dog",
		"the N ick broRn foS jumps oveB theTlazyEdQg",
		"the quWck b Ewn foxMjumNs Qver She lRzy dog",
	}
	if !reflect.DeepEqual(textCopies, wantedTextCopies) {
		test.Fail()
	}
}

func Test_search_withoutSample(test *testing.T) {
	variants := []string{"t***", "te**", "tes*"}
	sample := "test"
	result := search(variants, sample)
	if result != "tes*" {
		test.Fail()
	}
}

func Test_search_withSample(test *testing.T) {
	variants := []string{"t***", "te**", "test"}
	sample := "test"
	result := search(variants, sample)
	if result != "test" {
		test.Fail()
	}
}
