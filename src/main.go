package main

import (
	"runtime"
	"sync"
)

// Code is just translated from
// https://github.com/COSC3020/mergesort-tlaceby
func mergesort(arr []int) []int {
	n := len(arr)

	for size := 1; size < n; size *= 2 {
		for left := 0; left < n-size; left += 2 * size {
			mid := left + size - 1
			right := min(left+2*size-1, n-1)
			merge(arr, left, mid, right)
		}
	}

	return arr
}

func merge(arr []int, left, mid, right int) {
	leftIndex := left
	rightIndex := mid + 1

	// While there exist left or right subarrays
	for leftIndex <= mid && rightIndex <= right {
		// Do nothing as the left is less than right
		if arr[leftIndex] <= arr[rightIndex] {
			leftIndex++
			continue
		}

		// Since the first element in left subarray is > then first element in right subarray
		value := arr[rightIndex] // smallest of the two
		shiftIndex := rightIndex

		for shiftIndex != leftIndex {
			arr[shiftIndex] = arr[shiftIndex-1]
			shiftIndex--
		}

		arr[leftIndex] = value
		leftIndex++
		mid++
		rightIndex++
	}
}

func parallelMergesort(arr []int) {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Make sure go uses most cores possible
	var wg sync.WaitGroup
	var numGoroutines = runtime.NumCPU()
	var chunkSize = len(arr) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		var start = i * chunkSize
		var end = start + chunkSize

		if i == numGoroutines-1 {
			end = len(arr)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			// Creates a reference/window into the slice. DOES NOT create a copy
			mergesort(arr[start:end])
		}(start, end)
	}

	wg.Wait()

	// Finaly call mergesort on the final array. This means there will be
	mergesort(arr)
}

func main() {
	arr := []int{1, 4, 6, 21, 5, 2, 65, 2, 9}
	sorted := []int{1, 2, 2, 4, 5, 6, 9, 21, 65}

	parallelMergesort(arr)

	for indx, val := range arr {
		if sorted[indx] != val {
			panic("Array out of order")
		}
	}

	println("Passed")
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
