package main

import "testing"

func TestCanSubtractOneFromANumber(test *testing.T) {
	if integer := subtractByOne(2); integer != 1 {
		test.Error("The returned value was not what we expected")
	} else {
		test.Log("Can subtract a number from any int!")
	}
}
