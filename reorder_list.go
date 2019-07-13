package main

import (
	"fmt"
	"github.com/j3rrywan9/go-algorithms/linkedlist"
)

// Split a linked list in two halves
func splitList(head *linkedlist.Node) (*linkedlist.Node, *linkedlist.Node) {
	runner, walker := head, head

	for runner != nil && runner.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
	}

	middle := walker.Next
	walker.Next = nil

	return head, middle
}

func reverseList(head *linkedlist.Node) *linkedlist.Node {
	var dummy *linkedlist.Node

	for head != nil {
		current := head
		head = head.Next
		current.Next = dummy
		dummy = current
	}

	return dummy
}

// Merge two linked lists
func mergeLists(a, b *linkedlist.Node) *linkedlist.Node {
	head, tail := a, a

	a = a.Next

	for b != nil {
		tail.Next = b
		tail = tail.Next
		b = b.Next
		if a != nil {
			a, b = b, a
		}
	}

	return head
}

// LeetCode OJ No. 143
func reorderList(head *linkedlist.Node) *linkedlist.Node {
	if head == nil || head.Next == nil {
		return head
	}

	a, b := splitList(head)
	b = reverseList(b)
	head = mergeLists(a, b)

	return head
}

func main() {
	myList := &linkedlist.Node{1, nil}
	fmt.Println("Creating a 10-node linked list...")
	myList.Create(10)
	myList.Print()

	fmt.Println("Reordering above linked list...")
	head := reorderList(myList)
	head.Print()
}
