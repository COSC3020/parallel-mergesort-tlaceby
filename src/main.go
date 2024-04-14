package main

// Code is just translated from
// https://github.com/COSC3020/mergesort-tlaceby
func mergeSort(arr []int) []int {
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

func main() {
	arr := []int{1, 4, 6, 21, 5, 2, 65, 2, 9}
	sorted := []int{1, 2, 2, 4, 5, 6, 9, 21, 65}

	mergeSort(arr)

	for indx, val := range arr {
		if sorted[indx] != val {
			panic("Array out of order")
		}
	}

	println("Passed")
}
