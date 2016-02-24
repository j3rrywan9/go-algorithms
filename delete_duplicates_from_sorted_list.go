package main

import (
	"fmt"
	"github.com/j3rrywan9/go-algorithms/linkedlist"
)

// LeetCode OJ No. 83
func deleteDuplicates(head *linkedlist.Node) *linkedlist.Node {
	if head == nil {
		return head
	} else {
		current := head
		for current.Next != nil {
			if current.Val == current.Next.Val {
				current.Next = current.Next.Next
			} else {
				current = current.Next
			}
		}
		return head
	}
}

func main() {
	myList := &linkedlist.Node{1, nil}
	fmt.Println("Creating a 10-node sorted linked list...")
	myList.Create(10)
	myList.Print()
	fmt.Println("Adding a duplicate node to above linked list...")
	duplicateNode := &linkedlist.Node{2, nil}
	duplicateNode.Next = myList.Next.Next
	myList.Next.Next = duplicateNode
	myList.Print()

	fmt.Println("Deleting duplicate nodes in above linked list...")
	myList = deleteDuplicates(myList)
	myList.Print()
}
