package main

import "fmt"

// Stack structure using a slice
type Stack struct {
	items []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
// If the stack is empty, it returns -1 and a bool flag indicating that the stack was empty
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}

	// Get the last element (top of the stack)
	item := s.items[len(s.items)-1]
	// Remove the last element
	s.items = s.items[:len(s.items)-1]

	return item, true
}

// Top returns the top element of the stack without removing it
// If the stack is empty, it returns -1 and a bool flag indicating that the stack was empty
func (s *Stack) Top() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}
	return s.items[len(s.items)-1], true
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Empty returns true if the stack is empty, otherwise false
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func Swap(stack1, stack2 *Stack) {
	tempStack := stack1.items
	stack1.items = stack2.items
	stack2.items = tempStack
	return
}

func main() {
	// Create a new stack
	stack := Stack{}

	// Push elements onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Get the current size of the stack
	fmt.Println("Size:", stack.Size()) // Output: Size: 3

	// Peek the top element of the stack
	if top, ok := stack.Top(); ok {
		fmt.Println("Top element:", top) // Output: Top element: 30
	} else {
		fmt.Println("Stack is empty")
	}

	// Pop elements from the stack
	if popped, ok := stack.Pop(); ok {
		fmt.Println("Popped:", popped) // Output: Popped: 30
	} else {
		fmt.Println("Stack is empty")
	}

	// Check if the stack is empty
	fmt.Println("Is stack empty?", stack.Empty()) // Output: Is stack empty? false

	// Pop the remaining elements
	stack.Pop()
	stack.Pop()

	// Check if the stack is empty again
	fmt.Println("Is stack empty after popping all?", stack.Empty()) // Output: Is stack empty after popping all? true

	// Create two stacks
	stack1 := &Stack{}
	stack2 := &Stack{}

	// Add values to stack1
	stack1.Push(10)
	stack1.Push(20)
	stack1.Push(30)

	// Add values to stack2
	stack2.Push(40)
	stack2.Push(50)

	// Print stacks before swap
	fmt.Println("Stack 1 before swap:", stack1.items) // Output: [10 20 30]
	fmt.Println("Stack 2 before swap:", stack2.items) // Output: [40 50]

	// Swap stacks
	Swap(stack1, stack2)

	// Print stacks after swap
	fmt.Println("Stack 1 after swap:", stack1.items) // Output: [40 50]
	fmt.Println("Stack 2 after swap:", stack2.items) // Output: [10 20 30]
}
