package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"slices"
	"testing"
	"time"
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

func TestEmptyCase(t *testing.T) {
	arr := []int{}
	expected := []int{}

	parallelMergesort(arr)
	assert(arr, expected, t)
}

func TestSortingArrays(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Make sure go uses most cores possible
	var startTime = time.Now()
	var numberSlices = 100 // iterations

	for i := 0; i < numberSlices; i++ {
		var arrSize = randRange(1, 20000)
		var arr = make([]int, arrSize)
		var secondCopy = make([]int, arrSize)

		for index := range arr {
			arr[index] = rand.Int()
		}

		copy(secondCopy, arr)
		parallelMergesort(arr)  // parallel mergesort
		slices.Sort(secondCopy) // go's built-in sort for a easy check of valid sorting
		assert(secondCopy, arr, t)
	}

	var elapsedTime = time.Since(startTime).Seconds()
	fmt.Printf("All %d Tests Passed in %.1f seconds.\n", numberSlices, elapsedTime)
}
