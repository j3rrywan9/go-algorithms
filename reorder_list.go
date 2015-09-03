package main

import (
	"fmt"
)

type ListNode struct {
	val int
	next *ListNode
}

func createList(start, end int) *ListNode {
	head := &ListNode{start, nil}
	current := head
	for i := start + 1; i <= end; i++ {
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

func splitList(head *ListNode) (*ListNode, *ListNode) {
	runner, walker := head, head
	for runner != nil && runner.next != nil {
		walker = walker.next
		runner = runner.next.next
	}
	middle := walker.next
	walker.next = nil
	return head, middle
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

func mergeList(a, b *ListNode) *ListNode {
	head, tail := a, a
	a = a.next
	for b != nil {
		tail.next = b
		tail = tail.next
		b = b.next
		if a != nil {
			a, b = b, a
		}
	}
	return head
}

// LCOJ No. 143
func reorderList(pListNode *ListNode) *ListNode {
	if pListNode == nil || pListNode.next == nil {
		return pListNode
	}
	a, b := splitList(pListNode)
	b = reverseList(b)
	head := mergeList(a, b)
	return head
}

func main() {
	var l = createList(1, 4)
	printList(l)
	head := reorderList(l)
	printList(head)
}

