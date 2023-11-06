// 1. init array
// 2. insert data into an array (insert_at_last)
// 3. insert at first position
// 4. insert at n index position
// 5. remove an element from array at n index
// 7. update an element at index n
// 8. is having given an element and return index, else return -1
// Todo: must write time and space complexity
package main

import "fmt"

// ! initialization of an array
var arr []int

// Time complexity: O(1)
// Space complexity: O(1)
func insertAtLast(n int) {
	arr = append(arr, n)
}

// Time complexity: O(n) because copying all data to newArray requires n
// Space complexity: O(n) it is too
func insertAtFirst(n int) {
	var newArray []int
	newArray = append(newArray, n)
	newArray = append(newArray, arr...)
	arr = newArray
}

// Time complexity: O(n)
// Space complexity: O(n)
func insertAtIndexN(value int, index int) {
	if index < 0 {
		return
	}
	var newArray []int
	newArray = append(newArray, arr[:index]...)
	newArray = append(newArray, value)
	newArray = append(newArray, arr[index:]...)
	arr = newArray
}

// Time complexity: O(n)
// Space complexity: O(n)
func removeElementFromNPosition(index int) {
	if index < 0 {
		return
	}
	var newArray []int
	newArray = append(newArray, arr[:index]...)
	newArray = append(newArray, arr[index+1:]...)
	arr = newArray

}

// Time complexity: O(1)
// Space complexity: O(1)
func updateElementWithIndex(value, index int) {
	if index < 0 {
		return
	}
	arr[index] = value
}

// Time complexity: O(n)
// Space complexity: O(1)
func printArray() {
	fmt.Println("Array: ", arr)
}
func main() {
	insertAtLast(1)
	insertAtLast(2)
	insertAtLast(3)
	insertAtLast(4)
	insertAtLast(5)
	insertAtLast(6)
	insertAtLast(7)
	insertAtLast(8)
	insertAtFirst(0)
	insertAtIndexN(33, 0)
	printArray()
	removeElementFromNPosition(0)
	printArray()
}
