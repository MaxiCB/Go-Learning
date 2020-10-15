package main

import "fmt"

// OPERATIONS

// Node Type Struct
type Node struct {
	data string
	next *Node
}

// List Type Struct
type List struct {
	len int
	head *Node
}

func initializeList() *List {
	return &List{}
}

func (l *List) addFront(data string) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.len++
	return
}

func (l *List) addBack(data string) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	l.len++
	return
}

func (l *List) getFront() (string, error) {
	if l.head == nil {
		return "", fmt.Errorf("List is empty")
	}
	return l.head.data, nil
}

func (l *List) removeFront() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}
	l.head = l.head.next
	l.len--
	return nil
}

func (l *List) removeBack() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}
	var previous *Node
	current := l.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	if previous != nil {
		previous.next = nil
	} else {
		l.head = nil
	}
	l.len--
	return nil
}

func (l *List) traverse() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

func (l *List) size() int {
	return l.len
}

func main() {
	singleList := initializeList()
	fmt.Println("Adding to front.")
	singleList.addFront("1")
	fmt.Println("Adding to back.")
	singleList.addBack("2");
	fmt.Println("Adding to front.")
	singleList.addFront("3")
	fmt.Println("Adding to back.")
	singleList.addBack("4");
	singleList.traverse()
	fmt.Println(singleList.size())
}