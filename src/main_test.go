package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func assert(got []int, want []int, t *testing.T) {
	if len(got) != len(want) {
		t.Errorf("Slices dont have the same size: %d vs %d\n", len(want), len(got))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Slices Dont Match at index %d: (%d vs %d\n)", i, want[i], got[i])
		}
	}
}

func TestSortingArrays(t *testing.T) {
	var numberSlices = 100 // iterations

	for i := 0; i < numberSlices; i++ {
		var arrSize = randRange(1000, 25000)
		var arr = make([]int, arrSize)
		var arrCopy = make([]int, arrSize)
		var secondCopy = make([]int, arrSize)

		// Populate Array
		for index := range arr {
			arr[index] = rand.Int()
		}

		// Make sure to keep a copy to compare against
		copy(arrCopy, arr)
		copy(secondCopy, arr)

		parallelMergesort(arr)  // parallel mergesort
		mergesort(arrCopy)      // standard mergesort
		slices.Sort(secondCopy) // go's built-in sort

		assert(arrCopy, arr, t)
		assert(secondCopy, arr, t)
	}

	fmt.Printf("All %d Tests Passed\n", numberSlices)
}
