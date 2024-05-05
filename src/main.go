package main

import (
	"math"
	"runtime"
	"sync"
)

// Code inspired by my mergesort and async exersise.
// Additional resources include: https://www.geeksforgeeks.org/merge-sort-using-multi-threading/
func merge(arr []int, aux []int, left, mid, right int) {
	i, j, k := left, mid+1, left
	for k <= right {
		if i > mid {
			aux[k] = arr[j]
			j++
		} else if j > right {
			aux[k] = arr[i]
			i++
		} else if arr[i] <= arr[j] {
			aux[k] = arr[i]
			i++
		} else {
			aux[k] = arr[j]
			j++
		}
		k++
	}
	for k = left; k <= right; k++ {
		arr[k] = aux[k]
	}
}

func parallelMergeSort(arr, aux []int, left, right int, depthLimit int) {
	if left < right {
		if depthLimit == 0 {
			// depth is large enough do it sequentialy. We dont have infinite threads/cores
			mergesortSequential(arr, aux, left, right)
		} else {
			mid := (left + right) / 2
			var wg sync.WaitGroup
			wg.Add(2)

			// Seperate into two processes and wait for both to finish
			go func() {
				defer wg.Done()
				parallelMergeSort(arr, aux, left, mid, depthLimit-1)
			}()
			go func() {
				defer wg.Done()
				parallelMergeSort(arr, aux, mid+1, right, depthLimit-1)
			}()

			wg.Wait()
			merge(arr, aux, left, mid, right)
		}
	}
}

func mergesortSequential(arr, tmp_arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergesortSequential(arr, tmp_arr, left, mid)
		mergesortSequential(arr, tmp_arr, mid+1, right)
		merge(arr, tmp_arr, left, mid, right)
	}
}

func parallelMergesort(arr []int) {
	tmp_arr := make([]int, len(arr))
	depthLimit := int(math.Log2(float64(runtime.NumCPU()))) + 1
	parallelMergeSort(arr, tmp_arr, 0, len(arr)-1, depthLimit)
}

func main() {
	arr := []int{1, 4, 6, 21, 5, 2, 65, 2, 9}
	expected := []int{1, 2, 2, 4, 5, 6, 9, 21, 65}

	parallelMergesort(arr)

	for i, v := range arr {
		if expected[i] != v {
			panic("Array out of order")
		}
	}

	println("Passed")
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
