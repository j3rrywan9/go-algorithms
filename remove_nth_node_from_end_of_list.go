package main

import (
	"fmt"
	"github.com/j3rrywan9/go-algorithms/linkedlist"
)

// LeetCode OJ No. 19
func removeNthNodeFromEndOfList(head *linkedlist.Node, n int) *linkedlist.Node {
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// If n th node from end is the same as head, return head.Next directly
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	myList := &linkedlist.Node{Val: 1, Next: nil}
	fmt.Println("Creating a 10-node linked list...")
	myList.Create(10)
	myList.Print()

	fmt.Println("Removing 4th node from end of above linked list...")
	myList = removeNthNodeFromEndOfList(myList, 4)
	myList.Print()
}
