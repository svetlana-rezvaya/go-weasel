package weasel

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_mutate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	result := mutate(text, 0.2)

	wantedResult := "the qu Pk brown fox jumps oveF tGD Nazy dog"
	if result != wantedResult {
		test.Fail()
	}
}

func TestPopulate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	textCopies := Populate(text, 0.2, 3)

	wantedTextCopies := []string{
		"the qu Pk brown fox jumps oveF tGD Nazy dog",
		"the N ick broRn foS jumps oveB theTlazyEdQg",
		"the quWck b Ewn foxMjumNs Qver She lRzy dog",
	}
	if !reflect.DeepEqual(textCopies, wantedTextCopies) {
		test.Fail()
	}
}
