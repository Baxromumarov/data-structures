package main

import "fmt"

// 1. Bubble Sort O(n^2)
// 2. Selection Sort O(n^2)
// 3. Insertion Sort O(n^2)
// 4. Merge Sort O(n log n)
// 5. Quick Sort O(n^2) in the worst case, O(n log n) on average
// 6. Heap Sort O(n log n)
// 7. Counting Sort O(n + k), where k is the range of input
// 8. Radix Sort O(n * k), where k is the number of digits in the maximum number
// 9. Bucket Sort  O(n^2) in the worst case, O(n + k) on average
var nums = []int{12, 45, 19, 72, 34, 61, 83, 50, 27, 94, 16, 78, 39, 56, 87, 13, 75, 32, 69}

// Time complexity: O(n^2)
// Space complexity: O(1)
// Description: https://en.wikipedia.org/wiki/Bubble_sort#/media/File:Bubble-sort-example-300px.gif
func bubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// Time complexity: O(n^2)
// Space complexity: O(1)
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// Find the minimum element in the unsorted part of the array
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			fmt.Println(arr)
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the element at index i
		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}
}

// Time complexity: O(n^2)
// Space complexity: O(1)
// description: https://en.wikipedia.org/wiki/Insertion_sort#/media/File:Insertion-sort-example-300px.gif
func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// Time complexity: O(nlogn)
// Space complexity: O(n)
// description: https://en.wikipedia.org/wiki/Merge_sort#/media/File:Merge-sort-example-300px.gif
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	fmt.Println(left, right)

	// Merge the two sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)  // Append remaining elements from the left slice
	result = append(result, right[j:]...) // Append remaining elements from the right slice

	return result
}

// Time complexity: O(nlogn) on average
// Space complexity: O(1)
// description: https://en.wikipedia.org/wiki/Quicksort#/media/File:Sorting_quicksort_anim.gif
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for _, value := range arr[1:] {
		if value <= pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

// Time complexity:  O(n * log(n))
// Space complexity:  O(1)
// description: https://en.wikipedia.org/wiki/Heapsort#/media/File:Sorting_heapsort_anim.gif
func heapSort(arr []int) {
	n := len(arr)

	// Build a max-heap from the unsorted array
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i > 0; i-- {
		// Swap the root (maximum) element with the last element
		arr[0], arr[i] = arr[i], arr[0]

		// Call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Find the largest element among the root and its children
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not the root, swap them
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}
func main() {
	fmt.Println("Before sorted: ", nums)
	heapSort(nums)
	fmt.Println("After sorted: ", nums)

}
