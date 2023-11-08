// Implementation
package main

import "fmt"

// Singly linked lists
// Doubly linked lists
// Circular linked lists
// Circular doubly linked lists
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
