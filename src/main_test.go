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
	var pwon float32 = 0   // # times parallel won
	var mwon float32 = 0   // # times default merge won

	for i := 0; i < numberSlices; i++ {
		var arrSize = randRange(1, 20000)
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

		start := time.Now()
		parallelMergesort(arr) // parallel mergesort
		elapsed := time.Since(start)
		var best = elapsed

		start = time.Now()
		mergesort(arrCopy) // standard mergesort
		elapsed = time.Since(start)

		if best < elapsed {
			pwon++
		} else {
			mwon++
		}

		slices.Sort(secondCopy) // go's built-in sort for a easy check of valid sorting

		assert(arrCopy, arr, t)
		assert(secondCopy, arr, t)
	}

	var elapsedTime = time.Since(startTime).Seconds()
	fmt.Printf("All %d Tests Passed in %.1f seconds.\n", numberSlices, elapsedTime)
	fmt.Printf("Parallel Merge: %d \n", int(pwon))
	fmt.Printf("Standard Merge: %d \n", int(mwon))

	fmt.Printf("Parallel won %.1f%% \n", ((pwon / float32(numberSlices)) * 100))
}
