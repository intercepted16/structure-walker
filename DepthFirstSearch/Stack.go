package DepthFirstSearch

type Node struct {
	value [][]int
	next  *Node
}

type Stack struct {
	head *Node
	tail *Node
}

// createStack creates a new stack
func createStack(val [][]int) *Stack {
	node := &Node{value: val}
	stack := &Stack{head: node, tail: node}
	return stack
}

// push adds a new node to the stack
func (s *Stack) push(val [][]int) {
	node := &Node{value: val}
	if s.tail != nil {
		s.tail.next = node
	} else {
		s.head = node
	}
	s.tail = node
}

// pop removes the last node from the stack
func (s *Stack) pop() *[][]int {
	if s.head == nil {
		return nil
	}
	if s.head == s.tail {
		val := s.head.value
		s.head = nil
		s.tail = nil
		return &val
	}
	node := s.head
	for node.next != s.tail {
		node = node.next
	}
	val := s.tail.value
	s.tail = node
	s.tail.next = nil
	return &val
}

// peek returns the value of the last node in the stack
func (s *Stack) peek() [][]int {
	if s.tail == nil {
		return nil
	}
	return s.tail.value
}
