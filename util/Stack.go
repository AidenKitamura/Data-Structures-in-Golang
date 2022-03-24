package main

import (
	"fmt"
)

type Stacker interface {
	push()
	pop()
	getSize()
	isEmpty()
}

type StackElement struct {
	element         int
	previousElement *StackElement
}

type Stack struct {
	size           int
	currentElement *StackElement
}

func (s *Stack) push(e int) {
	if s.isEmpty() {
		s.currentElement = &StackElement{e, nil}
		s.size += 1
	} else {
		s.currentElement = &StackElement{e, s.currentElement}
		s.size += 1
	}
}

func (s *Stack) pop() (bool, int) {
	if s.isEmpty() {
		return false, 0
	} else {
		ele := s.currentElement.element
		s.currentElement = s.currentElement.previousElement
		s.size -= 1
		return true, ele
	}
}

func (s Stack) getSize() int {
	return s.size
}

func (s Stack) isEmpty() bool {
	if s.getSize() == 0 {
		return true
	} else {
		return false
	}
}

func (s Stack) String() string {
	str := "Current Elements: "
	for pos := s.currentElement; pos != nil; pos = pos.previousElement {
		str += fmt.Sprintf("%d, ", pos.element)
	}
	return str
}

func main() {
	fmt.Println("------Stack Testing------")
	s := Stack{}
	s.push(0)
	s.push(1)
	s.push(2)
	fmt.Println(s.String())
	for flag := s.isEmpty(); flag == false; flag = s.isEmpty() {
		_, ele := s.pop()
		fmt.Printf("Current Popped element: %d\n", ele)
	}
}
