package main

import (
	"fmt"
)

type Queue struct {
	stack1 []int
	stack2 []int
	size   int
}

func (q *Queue) Push(x int) {
	q.stack1 = append(q.stack1, x)
	q.size += 1
}

func (q *Queue) Pop() {
	if q.stack2 != nil {
		q.stack2 = q.stack2[:len(q.stack2)-1]
	} else {
		for i := 0; i < q.size-1; i++ {
			x := q.stack1[len(q.stack1)-1]
			q.stack1 = q.stack1[:len(q.stack1)-1]
			q.stack2 = append(q.stack2, x)
		}
		q.stack1 = q.stack1[:len(q.stack1)-1]
	}
	q.size -= 1
}

func (q *Queue) Peek() int {
	if q.stack2 != nil {
		return q.stack2[len(q.stack2)-1]
	} else {
		for i := 0; i < q.size; i++ {
			x := q.stack1[len(q.stack1)-1]
			q.stack1 = q.stack1[:len(q.stack1)-1]
			q.stack2 = append(q.stack2, x)
		}
		return q.stack2[len(q.stack2)-1]
	}
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Print() {
	fmt.Println("stack2:", q.stack2)
	fmt.Println("stack1:", q.stack1)
	fmt.Println("queue:", append(q.stack2, q.stack1...))
}

func main() {
	myqueue := new(Queue)
	myqueue.Push(1)
	myqueue.Push(2)
	myqueue.Push(3)
	myqueue.Push(4)
	myqueue.Push(5)
	myqueue.Pop()
	test := myqueue.Peek()
	fmt.Printf("%d\n", test)
	myqueue.Pop()
	myqueue.Print()
	fmt.Printf("Is queue empty? %t\n", myqueue.Empty())
}
