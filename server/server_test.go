package main_test

import "testing"

func TestTruth(t *testing.T) {
	if true == false {
		t.Error("true must not be equal to false. Madness!")
	}
}
