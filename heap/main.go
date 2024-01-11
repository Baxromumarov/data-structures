package main

import (
	"fmt"
)

// To heapify a subtree rooted with node i
// which is an index in arr[].
// n is size of heap
func heapify(arr []int, N, i int) {
	// Initialize largest as root
	largest := i

	// left = 2*i + 1
	l := 2*i + 1

	// right = 2*i + 2
	r := 2*i + 2

	// If left child is larger than root
	if l < N && arr[l] > arr[largest] {
		largest = l
	}

	// If right child is larger than largest so far
	if r < N && arr[r] > arr[largest] {
		largest = r
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		heapify(arr, N, largest)
	}
}

// Main function to do heap sort
func heapSort(arr []int, N int) {
	// Build heap (rearrange array)
	for i := N/2 - 1; i >= 0; i-- {
		heapify(arr, N, i)
	}

	// One by one extract an element from heap
	for i := N - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// A utility function to print array of size n
func printArray(arr []int, N int) {
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// Driver's code
func main() {
	arr := []int{12, 11, 13, 5, 6, 7, -3, 24, 32, 3, 0}
	N := len(arr)

	// Function call
	heapSort(arr, N)

	fmt.Println("Sorted array is")
	printArray(arr, N)
}
