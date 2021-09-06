package weasel

import "testing"

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

func TestSearch_withoutSample(test *testing.T) {
	variants := []string{"t***", "te**", "tes*"}
	sample := "test"
	result := Search(variants, sample)
	if result != "tes*" {
		test.Fail()
	}
}

func TestSearch_withSample(test *testing.T) {
	variants := []string{"t***", "te**", "test"}
	sample := "test"
	result := Search(variants, sample)
	if result != "test" {
		test.Fail()
	}
}
