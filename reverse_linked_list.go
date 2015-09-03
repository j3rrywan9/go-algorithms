package main

import (
	"fmt"
)

type ListNode struct {
	val int
	next *ListNode
}

func createList() *ListNode {
	head := &ListNode{1, nil}
	current := head
	for i := 2; i <= 10; i++ {
		pListNode := &ListNode{i, nil}
		current.next = pListNode
		current = pListNode
	}
	return head
}

func printList(pListNode *ListNode) {
	current := pListNode
	for {
		fmt.Printf("%d", current.val)
		if current.next != nil {
			fmt.Print("->")
			current = current.next
		} else {
			break
		}
	}
	fmt.Println()
}

func reverseList(pListNode *ListNode) *ListNode {
	current := pListNode
	var head *ListNode = nil
	for {
		if current == nil {
			break
		}
		temp := current.next
		current.next = head
		head = current
		current = temp
	}
	return head
}

func main() {
	var pListNode = createList()
	printList(pListNode)
	head := reverseList(pListNode)
	printList(head)
}

