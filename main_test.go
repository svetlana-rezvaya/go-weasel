package main

import "testing"

func Test_initialize(test *testing.T) {
	result := initialize(5)
	if result != "#####" {
		test.Fail()
	}
}
