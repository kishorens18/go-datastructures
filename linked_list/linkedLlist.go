package main

import "log"

type Node struct {
	data int
	next *Node
}

var head, tail *Node

func main() {
	for i := 0; i < 5; i++ {
		addValues(i)
	}

	addAtFirst(18)
	addAtLast(17)

	addAtRandomPos(4, 19)

	// Uncomment these lines to test deletion methods
	// deleteAtFirst()
	// deleteAtLast()

	deleteAtRandomPos(4)

	printList() // This prints the entire list without losing reference to head

	NodesEvenOrOdd()
}

// Print the linked list
func printList() {
	tempNode := head
	for tempNode != nil {
		println(tempNode.data)
		tempNode = tempNode.next
	}
}

// Delete the first node
func deleteAtFirst() {
	if head != nil {
		head = head.next
		if head == nil { // If list becomes empty, reset tail as well
			tail = nil
		}
	}
}

// Delete the node at a specific position
func deleteAtRandomPos(position int) {
	if position == 0 {
		deleteAtFirst()
		return
	}
	tempNode := head
	for i := 0; tempNode != nil && i < position-1; i++ {
		tempNode = tempNode.next
	}

	if tempNode == nil || tempNode.next == nil {
		log.Println("Position out of bounds")
		return
	}

	tempNode.next = tempNode.next.next
	if tempNode.next == nil { // If deleting the last element, update tail
		tail = tempNode
	}
}

// Delete the last node
func deleteAtLast() {
	if head == nil {
		return
	}
	if head.next == nil {
		head, tail = nil, nil
		return
	}

	tempNode := head
	for tempNode.next.next != nil {
		tempNode = tempNode.next
	}
	tempNode.next = nil
	tail = tempNode
}

// Add a node at a specific position
func addAtRandomPos(position int, value int) {
	if position == 0 {
		addAtFirst(value)
		return
	}

	tempNode := head
	for i := 0; tempNode != nil && i < position-1; i++ {
		tempNode = tempNode.next
	}

	if tempNode == nil {
		log.Println("Position out of bounds")
		return
	}

	newNode := createNode(value)
	newNode.next = tempNode.next
	tempNode.next = newNode
	if newNode.next == nil {
		tail = newNode
	}
}

// Add a node at the beginning
func addAtFirst(val int) {
	newNode := createNode(val)
	newNode.next = head
	head = newNode
	if tail == nil {
		tail = head
	}
}

// Add a node at the end
func addAtLast(val int) {
	newNode := createNode(val)
	if head == nil {
		head = newNode
		tail = head
		return
	}
	tail.next = newNode
	tail = newNode
}

// Add a new node to the end of the list
func addValues(value int) {
	newNode := createNode(value)
	if head == nil {
		head = newNode
		tail = head
		return
	}
	tail.next = newNode
	tail = newNode
}

// Helper function to create a new node
func createNode(val int) *Node {
	return &Node{
		data: val,
	}
}

func NodesEvenOrOdd() {
	size := false

	if head == nil {
		log.Println("List is Empty...")
	}

	for head.next != nil {
		size = !size
		head = head.next
	}

	if size {
		println("List Is Even...")
	} else {
		println("List Is Odd...")
	}
}
