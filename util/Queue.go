package main

import (
	"fmt"
)

type Queueer interface {
	push()
	pop()
	getSize()
	isEmpty()
}

type QueueElement struct {
	element     int
	nextElement *QueueElement
}

type Queue struct {
	size           int
	currentElement *QueueElement
	lastElement    *QueueElement
}

func (q *Queue) push(e int) {
	if q.isEmpty() {
		q.currentElement = &QueueElement{e, nil}
		q.lastElement = q.currentElement
		q.size += 1
	} else {
		q.lastElement.nextElement = &QueueElement{e, nil}
		q.lastElement = q.lastElement.nextElement
		q.size += 1
	}
}

func (q *Queue) pop() (bool, int) {
	if q.isEmpty() {
		return false, 0
	} else {
		ele := q.currentElement.element
		if q.getSize() == 1 {
			q.currentElement = nil
			q.lastElement = nil
		} else {
			q.currentElement = q.currentElement.nextElement
		}
		q.size -= 1
		return true, ele
	}
}

func (q Queue) getSize() int {
	return q.size
}

func (q Queue) isEmpty() bool {
	if q.getSize() == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) String() string {
	str := "Current Elements: "
	for pos := q.currentElement; pos != nil; pos = pos.nextElement {
		str += fmt.Sprintf("%d, ", pos.element)
	}
	return str
}

func main() {
	fmt.Println("------Queue Testing------")
	s := Queue{}
	s.push(0)
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.String())
	for flag := s.isEmpty(); flag == false; flag = s.isEmpty() {
		_, ele := s.pop()
		fmt.Printf("Current Popped element: %d\n", ele)
	}
	flag, ele := s.pop()
	fmt.Printf("When popping empty Queue, we get: %v and %d", flag, ele)
}
