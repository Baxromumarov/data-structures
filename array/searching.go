// 1. Linear search
// 2. Binary search
// 3. Jump search
// 4. Interpolation Search (for sorted arrays with uniformly distributed values)
// 5. Exponential Search (for sorted arrays) (also called doubling search or galloping search or Struzik search)
// 6. Hashing
// 7. Fibonacci Search (for sorted arrays)
// 8. Ternary Search (for sorted arrays)
package main

import (
	"fmt"
	"math"
	"sort"
)

var Array = []int{12, 45, 19, 72, 34, 61, 83, 50, 27, 94, 16, 78, 39, 56, 87, 13, 75, 32, 69}

// Time complexity: O(n)
// Space complexity: O(1)
// todo: description in order, element check
func linearSearch(arr []int, element int) int {

	for index, val := range arr {
		if val == element {
			return index // +1
		}
	}

	// if not found an element
	return -1
}

// Time complexity: O(log(n))
// Space complexity: O(1)
// for sorted array
// Todo: description link: https://en.wikipedia.org/wiki/Binary_search_algorithm#/media/File:Binary-search-work.gif
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low +high)/2
		

		if arr[mid] == target {
			return mid // Element found at index mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Element is not present in the array
}

// Time complexity: O(√n)
// Space complexity: O(1)
// for sorted array
// Todo: description link: https://www.youtube.com/watch?v=63kS6ZkMpkA
func jumpSearch(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	prev := 0

	for arr[minNum(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))

		if prev >= n {
			return -1 // Element is not present in the array
		}
	}

	for arr[prev] < target {
		prev++

		if prev == minNum(step, n) {
			return -1 // Element is not present in the array
		}
	}

	if arr[prev] == target {
		return prev // Element found at index prev
	}

	return -1 // Element is not present in the array
}

//! linear search  <  jump search  <  binary search

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Time complexity: O(log log(n))
// Space complexity: O(1)
// for sorted array
// Todo: description link: https://en.wikipedia.org/wiki/Interpolation_search#/media/File:Interpolation_sort.gif
func interpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + ((target-arr[low])*(high-low))/(arr[high]-arr[low])

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Time complexity: O(log n)
// Space complexity: O(1)
// Exponential Search (for sorted arrays)
func exponentialSearch(arr []int, target int) int {
	n := len(arr)

	// If the target is the first element, return index 0
	if arr[0] == target {
		return 0
	}

	// Find the range for binary search
	i := 1
	for i < n && arr[i] <= target {
		i *= 2
	}

	// Perform binary search in the identified range
	return binarySearchForExponentialSearch(arr, target, i/2, minNum(i, n-1))
}

func binarySearchForExponentialSearch(arr []int, target int, low, high int) int {
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Element found at index mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Element is not present in the range
}

// Time complexity: O(log n)
// Space complexity: O(1)
func fibonacciSearch(arr []int, target int) int {
	n := len(arr)

	// Initialize Fibonacci numbers
	fibM2 := 0
	fibM1 := 1
	fib := fibM2 + fibM1

	// Find the smallest Fibonacci number that is greater than or equal to n
	for fib < n {
		fibM2 = fibM1
		fibM1 = fib
		fib = fibM2 + fibM1
	}

	// Initialize variables for the offset, i, and mid
	offset := -1

	for fibM2 > 1 {
		// Calculate the index
		i := min(offset+fibM2, n-1)

		// If the element at index i is greater than the target, move to the smaller subarray
		if arr[i] > target {
			fib = fibM2
			fibM1 = fibM1 - fibM2
			fibM2 = fib - fibM1
		} else if arr[i] < target { // If the element at index i is smaller than the target, move to the larger subarray
			fib = fibM1
			fibM1 = fibM2
			fibM2 = fib - fibM1
			offset = i
		} else { // Element found at index i
			return i
		}
	}

	// Check if the last remaining element is the target
	if fibM1 == 1 && arr[offset+1] == target {
		return offset + 1
	}

	// Element not found in the array
	return -1

}

// Time complexity: O(log3(n))
// Space complexity: O(1)
func ternarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		// Calculate partition points
		mid1 := low + (high-low)/3
		mid2 := high - (high-low)/3

		if arr[mid1] == target {
			return mid1
		} else if arr[mid2] == target {
			return mid2
		}

		if target < arr[mid1] {
			high = mid1 - 1
		} else if target > arr[mid2] {
			low = mid2 + 1
		} else {
			low = mid1 + 1
			high = mid2 - 1
		}
	}

	return -1
}

func main() {
	sort.Ints(Array)
	fmt.Println(Array)
	fmt.Println(ternarySearch(Array, 50))
}
