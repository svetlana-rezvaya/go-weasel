package weasel

import (
	"math/rand"
	"testing"
)

func TestInitialize(test *testing.T) {
	rand.Seed(1)

	result := Initialize(5)
	if result != "OPCLE" {
		test.Fail()
	}
}
