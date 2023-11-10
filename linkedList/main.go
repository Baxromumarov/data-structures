// Implementation
package main

import "fmt"

// 1. Singly linked lists
// 2. Doubly linked lists
// 3. Circular linked lists
// 4. Circular doubly linked lists
// ! Doubly linked lists
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
}

// Append adds a new node to the end of the doubly linked list
func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}
	if dll.Head == nil {
		dll.Head = newNode
	} else {
		current := dll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		newNode.Prev = current
	}
}
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

// DisplayBackward prints the elements of the doubly linked list in backward direction
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.Head
	for current != nil && current.Next != nil {
		current = current.Next
	}
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Prev
	}
	fmt.Println()
}
func main() {
	var dll DoublyLinkedList

	// Appending nodes
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)
	dll.Append(5)
	dll.Append(6)
	dll.Append(7)
	dll.Append(8)

	// Displaying in both directions
	fmt.Print("Forward: ")
	dll.DisplayForward()

	fmt.Print("Backward: ")
	dll.DisplayBackward()
}

//! Singly linked lists
/*
type SinglyLinkedList struct {
	Value int
	Next  *SinglyLinkedList
}

func (s *SinglyLinkedList) InsertAtEnd(val int) {
	newNode := &SinglyLinkedList{Value: val, Next: nil}

	if s.Next == nil {
		s.Next = newNode
	} else {
		s.Next.InsertAtEnd(val)
	}
}

func (s *SinglyLinkedList) Print() {
	current := s
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (s *SinglyLinkedList) Len() int {
	length := 0

	for head := s; head != nil; head = head.Next {
		length++
	}
	return length
}

func (s *SinglyLinkedList) InsertInNIndex(index, value int) {
	if index < 0 {
		// Invalid index, do nothing
		return
	}

	newNode := &SinglyLinkedList{Value: value, Next: nil}

	if index == 0 {
		// Insert at the beginning
		newNode.Next = s
		*s = *newNode
		return
	}

	current := s
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		// Index is out of bounds, do nothing
		return
	}

	newNode.Next = current.Next
	current.Next = newNode
}

func (s *SinglyLinkedList) DeleteFromNIndex(index int) {
	if index < 0 || s == nil {
		// Invalid index or empty list, do nothing
		return
	}

	if index == 0 {
		// Delete the first node
		*s = *s.Next
		return
	}

	current := s
	for i := 0; i < index-1 && current.Next != nil; i++ {
		current = current.Next
	}

	if current.Next == nil {
		// Index is out of bounds, do nothing
		return
	}

	current.Next = current.Next.Next

}

func main() {
	// Create the head node
	linkedList := &SinglyLinkedList{Value: 1}
	linkedList.InsertAtEnd(2)
	linkedList.InsertAtEnd(3)
	linkedList.InsertAtEnd(4)
	linkedList.InsertAtEnd(5)
	linkedList.InsertInNIndex(2, -19)
	fmt.Println("N: ", linkedList.Len())
	linkedList.DeleteFromNIndex(3)
	linkedList.Print()

}
*/
