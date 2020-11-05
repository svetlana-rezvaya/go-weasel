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
	result := mutate("test", 0)
	if result != "test" {
		test.Fail()
	}
}
