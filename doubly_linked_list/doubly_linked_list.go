package main

import "fmt"

// OPERATIONS

// Node Type Struct
type Node struct {
	data string
	next *Node
	prev *Node
}

// List Type Struct
type List struct {
	len int
	head *Node
	tail *Node
}

func initializeList() *List {
	return &List{}
}

func (l *List) addFront(data string) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.head
		l.head.next = node
		l.head = node
	}
	l.len++
	return
}

func (l *List) addBack(data string) {
	node := &Node{data: data}
	if l.head == nil || l.tail == nil{
		l.head = node
		l.tail = node
	} else {
		current := l.head
		for current.prev != nil {
			current = current.prev
		}
		current.prev = node
		node.next = current
		l.tail = node
	}
	l.len++
	return
}

func (l *List) addAfter(node *Node, data string) {
	new := &Node{data: data}
	new.next = node
	new.prev = node.prev
	node.prev.next = new
	node.prev = new

}


// [] [] []
// [] <> [] []
// [] <> [] []
func (l *List) addBefore(node *Node, data string) {
	new := &Node{data: data}
	new.prev = node
	new.next = node.next
	node.next.prev = new
	node.next = new
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
	l.head = l.head.prev
	l.len--
	return nil
}

func (l *List) removeBack() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}
	var previous *Node
	current := l.head
	for current.prev != nil {
		previous = current
		current = current.prev
	}
	if previous != nil {
		previous.prev = nil
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
		current = current.prev
	}
	return nil
}

func (l *List) traverseReversed() error {
	if l.tail == nil {
		return fmt.Errorf("List is empty")
	}
	current := l.tail
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
	doubleList := initializeList()
	doubleList.addFront("1")
	doubleList.addFront("2");
	doubleList.addFront("3")
	doubleList.addFront("4");
	head := doubleList.head
	tail := doubleList.tail
	doubleList.addAfter(head, "New")
	doubleList.addBefore(tail, "New Item")
	doubleList.traverse()
}