package linkedlist

import (
	"fmt"
)

// Node of a linked list
type Node struct {
	Val  int
	Next *Node
}

// Create a linked list of specified length
func (this *Node) Create(length int) {
	current := this
	for i := 2; i <= length; i++ {
		pNode := &Node{i, nil}
		current.Next = pNode
		current = pNode
	}
}

// Print the value of each node of a linked list
func (this *Node) Print() {
	current := this
	for {
		fmt.Printf("%d", current.Val)

		if current.Next != nil {
			fmt.Print("->")
			current = current.Next
		} else {
			break
		}
	}
	fmt.Println()
}
