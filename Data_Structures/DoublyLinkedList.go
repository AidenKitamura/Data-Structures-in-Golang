package main

import (
	"fmt"
)

type DoublyLinkedLister interface {
	add()
	get()
	set()
	pop()
	popFirst()
	getSize()
	remove()
	isEmpty()
}

type LinkedElement struct {
	element         int
	previousElement *LinkedElement
	nextElement     *LinkedElement
}

type DoublyLinkedList struct {
	size         int
	firstElement *LinkedElement
	lastElement  *LinkedElement
}

func (d *DoublyLinkedList) add(e int) {
	ele := LinkedElement{}
	ele.element = e
	if d.size == 0 {
		d.firstElement = &ele
		d.lastElement = &ele
		d.size += 1
	} else {
		d.lastElement.nextElement = &ele
		ele.previousElement = d.lastElement
		d.lastElement = &ele
		d.size += 1
	}
}

func (d *DoublyLinkedList) get(i int) int {
	if i >= d.getSize() || i < 0 {
		fmt.Printf("Index Out of Range. Cannot fetch element of index %d\n", i)
		return 0
	} else if i > d.getSize()/2 {
		currentElement := d.lastElement
		for pos := d.getSize() - 1; pos > i; pos-- {
			currentElement = currentElement.previousElement
		}
		return currentElement.element
	} else {
		currentElement := d.firstElement
		for pos := 0; pos < i; pos++ {
			currentElement = currentElement.nextElement
		}
		return currentElement.element
	}
}

func (d *DoublyLinkedList) set(i, e int) {
	if i >= d.getSize() || i < 0 {
		fmt.Printf("Index Out of Range. Cannot set element of index %d\n", i)
	} else if i > d.getSize()/2 {
		currentElement := d.lastElement
		for pos := d.getSize() - 1; pos > i; pos-- {
			currentElement = currentElement.previousElement
		}
		currentElement.element = e
	} else {
		currentElement := d.firstElement
		for pos := 0; pos < i; pos++ {
			currentElement = currentElement.nextElement
		}
		currentElement.element = e
	}
}

func (d *DoublyLinkedList) pop() int {
	if d.isEmpty() {
		fmt.Println("The list is empty, cannot pop.")
		return 0
	} else if d.size == 1 {
		ele := d.lastElement.element
		d.firstElement = nil
		d.lastElement = nil
		d.size -= 1
		return ele
	} else {
		ele := d.lastElement.element
		d.lastElement.previousElement.nextElement = nil
		d.lastElement = d.lastElement.previousElement
		d.size -= 1
		return ele
	}
}

func (d *DoublyLinkedList) popFirst() int {
	if d.isEmpty() {
		fmt.Println("The list is empty, cannot pop.")
		return 0
	} else if d.size == 1 {
		ele := d.firstElement.element
		d.firstElement = nil
		d.lastElement = nil
		d.size -= 1
		return ele
	} else {
		ele := d.firstElement.element
		d.firstElement.nextElement.previousElement = nil
		d.firstElement = d.firstElement.nextElement
		d.size -= 1
		return ele
	}
}

func (d *DoublyLinkedList) remove(i int) {
	if i >= d.getSize() || i < 0 {
		fmt.Printf("Index Out of Range. Cannot remove element of index %d\n", i)
	} else if d.getSize() == 1 {
		d.firstElement = nil
		d.lastElement = nil
		d.size -= 1
	} else if i > d.getSize()/2 {
		currentElement := d.lastElement
		for pos := d.getSize() - 1; pos > i; pos-- {
			currentElement = currentElement.previousElement
		}
		if i != d.getSize()-1 {
			currentElement.nextElement.previousElement = currentElement.previousElement
			currentElement.previousElement.nextElement = currentElement.nextElement
		} else {
			currentElement.previousElement.nextElement = nil
			d.lastElement = currentElement.previousElement
		}
		d.size -= 1
	} else {
		currentElement := d.firstElement
		for pos := 0; pos < i; pos++ {
			currentElement = currentElement.nextElement
		}
		if i != 0 {
			currentElement.previousElement.nextElement = currentElement.nextElement
			currentElement.nextElement.previousElement = currentElement.previousElement
		} else {
			currentElement.nextElement.previousElement = nil
			d.firstElement = currentElement.nextElement
		}
		d.size -= 1
	}
}

func (d *DoublyLinkedList) isEmpty() bool {
	if d.size == 0 {
		return true
	} else {
		return false
	}
}

func (d *DoublyLinkedList) getSize() int {
	return d.size
}

func (d *DoublyLinkedList) String() string {
	if d.isEmpty() {
		return "Currently the list has nothing with it......\n"
	} else {
		listContent := "Current elements: "
		for ptr := d.firstElement; ptr != nil; ptr = ptr.nextElement {
			listContent += fmt.Sprintf("%d, ", ptr.element)
		}
		listContent += "\n"
		return listContent
	}
}

func main() {
	ls := DoublyLinkedList{}
	ls.add(1)
	ls.add(2)
	ls.add(3)
	ls.add(4)
	fmt.Println(ls.String())
	ls.remove(6)
	ls.remove(2)
	ls.remove(0)
	fmt.Println(ls.String())
	ls.remove(0)
	ls.remove(0)
	fmt.Println(ls.String())
	ls.add(0)
	fmt.Printf("Popped: %d\n", ls.pop())
	ls.add(0)
	fmt.Printf("Popped First: %d\n", ls.popFirst())
	ls.pop()
}
