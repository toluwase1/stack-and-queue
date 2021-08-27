package main

import "fmt"

//ordered
//struct containing slices
type Stack struct {
	items []int
}

type Queue struct {
	items []int
}

//Push Method for stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop Method for stack
func (s *Stack) Pop() int {
	removable := len(s.items) - 1
	s.items = s.items[:removable]
	return removable
}

//Check if Stack is Empty
func (e *Stack) IsEmpty() bool {
	if len(e.items) == 0 {
		return true
	}
	return false
}
func (pe *Stack) Peek() int {
	value := pe.items[len(pe.items)-1]
	return value
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	removable := q.items[0]
	q.items = q.items[1:]
	return removable
}

func (p *Queue) Peek() int {
	value := p.items[0]
	return value
}

func main() {
	//myStack:= Stack{}
	//fmt.Println(myStack)
	//myStack.Push(10)
	//myStack.Push(20)
	//myStack.Push(30)
	//fmt.Println(myStack)
	//myStack.Pop()
	//fmt.Println(myStack)
	fmt.Println("------queues-------")
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(2000)
	myQueue.Enqueue(30000)
	fmt.Println(myQueue)
	myQueue.Dequeue()

	fmt.Println(myQueue)
	myStack := Stack{}
	myStack.Push(50)
	myStack.Push(20)
	fmt.Println(myStack)
	fmt.Println(myStack.IsEmpty())
	fmt.Println(myStack.Peek())
	fmt.Println(myQueue.Peek())

	a1 := &Queue{}
	a2 := &Queue{}

	a := &a1
	b := &a2
	fmt.Println(a)
	fmt.Println(b)

	//fmt.Println(*a1)
	//fmt.Println(a2)

	//b:=2
	//a:=&b
	//fmt.Println(a)

}

func IsValidParentheis(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 0 {
		return true
	}
	parentMap := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	var stack []rune

	for _, j := range s {
		if j == '(' || j == '{' || j == '[' {
			stack = append(stack, parentMap[j])
		} else if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != j {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return true && len(stack) == 0
}
