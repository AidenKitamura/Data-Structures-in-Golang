package main

import (
	"fmt"
)

type PriorityQueueer interface {
	push()
	pop()
	getSize()
	isEmpty()
}

type PriorityElement struct {
	element     int
	nextElement *PriorityElement
}

type PriorityQueue struct {
	size           int
	currentElement *PriorityElement
}

// Smaller Number has a higher Priority
func (p *PriorityQueue) push(e int) {
	if p.isEmpty() {
		p.currentElement = &PriorityElement{e, nil}
		p.size += 1
	} else {
		var prevPos *PriorityElement
		for pos := p.currentElement; pos != nil; pos = pos.nextElement {
			if pos.nextElement == nil {
				// Last element, must insert
				if prevPos == nil {
					// Only one element in queue
					if e <= pos.element {
						p.currentElement = &PriorityElement{e, pos}
					} else {
						pos.nextElement = &PriorityElement{e, nil}
					}
				} else {
					// Having Previous elements
					if e <= pos.element {
						prevPos.nextElement = &PriorityElement{e, pos}
					} else {
						pos.nextElement = &PriorityElement{e, nil}
					}
				}
				p.size += 1
				break
			} else {
				// Might be in between
				if prevPos == nil {
					// Being first element
					if e <= pos.element {
						p.currentElement = &PriorityElement{e, pos}
						p.size += 1
						break
					} else {
						prevPos = pos
						continue
					}
				} else {
					// Truly in between
					if e <= pos.element {
						prevPos.nextElement = &PriorityElement{e, pos}
						p.size += 1
						break
					} else {
						prevPos = pos
						continue
					}
				}
			}
		}
	}
}

func (p *PriorityQueue) pop() (bool, int) {
	if p.isEmpty() {
		return false, 0
	} else {
		element := p.currentElement.element
		p.currentElement = p.currentElement.nextElement
		p.size -= 1
		return true, element
	}
}

func (p PriorityQueue) getSize() int {
	return p.size
}

func (p PriorityQueue) isEmpty() bool {
	if p.getSize() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("------Priority Queue Testing------")
	p := PriorityQueue{}
	p.push(1)
	p.push(3)
	p.push(2)
	p.push(0)
	p.push(100)
	p.push(99)
	for flag := p.isEmpty(); flag == false; flag = p.isEmpty() {
		_, ele := p.pop()
		fmt.Println(ele)
	}
}
