package main

import (
	"fmt"
	"github.com/j3rrywan9/go-algorithms/linkedlist"
)

// LeetCode OJ No. 203
func removeElements(head *linkedlist.Node, val int) *linkedlist.Node {
	// Create a dummy node and link it before head
	dummyNode := &linkedlist.Node{-1, nil}
	dummyNode.Next = head
	current := dummyNode

	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummyNode.Next
}

func main() {
	myList := &linkedlist.Node{1, nil}
	fmt.Println("Creating a 10-node linked list...")
	myList.Create(10)
	myList.Print()

	fmt.Println("Removing all elements that have value 6 from above linked list...")
	myList = removeElements(myList, 6)
	myList.Print()
}
