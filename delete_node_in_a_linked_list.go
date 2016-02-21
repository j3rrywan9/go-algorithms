package main

import (
	"fmt"
	"github.com/j3rrywan9/go-algorithms/linkedlist"
)

// LeetCode OJ No. 237
func deleteNode(node *linkedlist.Node) {
	*node = *node.Next
}

func main() {
	myList := &linkedlist.Node{1, nil}
	fmt.Println("Creating a 10-node linked list...")
	myList.Create(10)
	myList.Print()

	fmt.Println("Deleting 3rd node from above linked list...")
	deleteNode(myList.Next.Next)
	myList.Print()
}
