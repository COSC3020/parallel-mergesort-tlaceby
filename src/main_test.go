package main

import (
	"math/rand/v2"
	"testing"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func assert(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEmptyArray(t *testing.T) {

}
