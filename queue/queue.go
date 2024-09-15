package main

import "fmt"

// Queue structure using a slice
type Queue struct {
	items []int
}

// Push adds an element to the back of the queue
func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
}

// Pop removes and returns the element from the front of the queue
// If the queue is empty, it returns -1 and a bool flag indicating the queue was empty
func (q *Queue) Pop() (int, bool) {
	if len(q.items) == 0 {
		return -1, false
	}

	// Get the first element
	item := q.items[0]
	// Remove the first element from the slice
	q.items = q.items[1:]

	return item, true
}

// Front returns the front element without removing it
// If the queue is empty, it returns -1 and a bool flag indicating the queue was empty
func (q *Queue) Front() (int, bool) {
	if len(q.items) == 0 {
		return -1, false
	}
	return q.items[0], true
}

// Back returns the last element without removing it
// If the queue is empty, it returns -1 and a bool flag indicating the queue was empty
func (q *Queue) Back() (int, bool) {
	if len(q.items) == 0 {
		return -1, false
	}
	return q.items[len(q.items)-1], true
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Empty returns true if the queue is empty, otherwise false
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

// Swap swaps the elements between two queues
func Swap(q1, q2 *Queue) {
	tempItems := q1.items
	q1.items = q2.items
	q2.items = tempItems
}

func main() {
	// Create two queues
	queue1 := &Queue{}
	queue2 := &Queue{}

	// Push elements onto queue1
	queue1.Push(10)
	queue1.Push(20)
	queue1.Push(30)

	// Push elements onto queue2
	queue2.Push(40)
	queue2.Push(50)

	// Print the front and back of queue1
	if front, ok := queue1.Front(); ok {
		fmt.Println("Front of queue1:", front) // Output: Front of queue1: 10
	}
	if back, ok := queue1.Back(); ok {
		fmt.Println("Back of queue1:", back) // Output: Back of queue1: 30
	}

	// Print size of queue1 and queue2 before swap
	fmt.Println("Size of queue1 before swap:", queue1.Size()) // Output: Size of queue1 before swap: 3
	fmt.Println("Size of queue2 before swap:", queue2.Size()) // Output: Size of queue2 before swap: 2

	// Swap queue1 and queue2
	Swap(queue1, queue2)

	// Print size of queue1 and queue2 after swap
	fmt.Println("Size of queue1 after swap:", queue1.Size()) // Output: Size of queue1 after swap: 2
	fmt.Println("Size of queue2 after swap:", queue2.Size()) // Output: Size of queue2 after swap: 3

	// Pop elements from queue1 (which was swapped with queue2)
	if popped, ok := queue1.Pop(); ok {
		fmt.Println("Popped from queue1:", popped) // Output: Popped from queue1: 40
	}

	// Check if queue1 is empty
	fmt.Println("Is queue1 empty?", queue1.Empty()) // Output: Is queue1 empty? false
}
