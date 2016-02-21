package main

import (
	"fmt"
	"github.com/j3rrywan9/go-algorithms/linkedlist"
)

// LeetCode OJ No. 206
func reverseLinkedList(head *linkedlist.Node) *linkedlist.Node {
	var dummy *linkedlist.Node
	for head != nil {
		current := head
		head = head.Next
		current.Next = dummy
		dummy = current
	}
	return dummy
}

func main() {
	myList := &linkedlist.Node{1, nil}
	fmt.Println("Creating a 10-node linked list...")
	myList.Create(10)
	myList.Print()

	fmt.Println("Reversing above linked list...")
	reversedList := reverseLinkedList(myList)
	reversedList.Print()
}
