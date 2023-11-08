// Implementation of the Stack
package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Item   []interface{}
	Length int
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.Item[s.Length-1], nil
}

func (s *Stack) Len() int {
	return s.Length
}

func (s *Stack) IsEmpty() bool {
	if len(s.Item) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Push(item interface{}) {
	s.Item = append(s.Item, item)
	s.Length++
}

func (s *Stack) Pop() {
	s.Item = s.Item[:s.Length-1]
	s.Length--
}

func (s *Stack) Print() {
	for i := 0; i < s.Length; i++ {
		if i == s.Length-1 {
			fmt.Print(" ", s.Item[i])
		} else {
			fmt.Printf("%v -> ", s.Item[i])
		}
	}
	fmt.Println()
}
func main() {
	var stc Stack

	stc.Push(12)
	stc.Push(113)
	stc.Push(145)
	stc.Push(98643)
	stc.Push(3534)
	stc.Push(983)
	stc.Push(5867)
	stc.Print()
	stc.Pop()
	stc.Print()

}
