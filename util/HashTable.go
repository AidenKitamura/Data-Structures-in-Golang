package util

import (
	"fmt"
)

type HashTableer interface {
	get()
	set()
	remove()
	getSize()
	isEmpty()
}

type HashTable struct {
	size     int
	elements [10]*HashElement
}

type HashElement struct {
	key         int
	element     string
	nextElement *HashElement
}

func hash(key int) int {
	// You can implement your own hashing function here
	return (key*key + 3) % 10
}

func (h HashTable) getSize() int {
	return h.size
}

func (h HashTable) isEmpty() bool {
	return h.size == 0
}

func (h HashTable) get(key int) string {
	for ptr := h.elements[hash(key)]; ptr != nil; ptr = ptr.nextElement {
		if ptr.key == key {
			return ptr.element
		}
	}
	fmt.Printf("Element of key %d does not exist.\n", key)
	return ""
}

func (h *HashTable) set(key int, ele string) {
	var dst *HashElement
	for ptr := h.elements[hash(key)]; ptr != nil; ptr = ptr.nextElement {
		if ptr.key == key {
			ptr.element = ele
			return
		}
		dst = ptr
	}

	if dst == nil {
		h.elements[hash(key)] = &HashElement{key, ele, nil}
	} else {
		dst.nextElement = &HashElement{key, ele, nil}
	}
	h.size += 1
}

func (h *HashTable) remove(key int) {
	var prev *HashElement
	for ptr := h.elements[hash(key)]; ptr != nil; ptr = ptr.nextElement {
		if ptr.key == key {
			if prev == nil {
				h.elements[hash(key)] = ptr.nextElement
			} else {
				prev.nextElement = ptr.nextElement
			}
			h.size -= 1
			return
		}
		prev = ptr
	}
	fmt.Printf("Element with key %d does not exist, cannot remove.\n", key)
}

func main() {
	fmt.Println("------Hash Table Testing------")
	ht := HashTable{}
	ht.set(0, "Char 0")
	ht.set(13, "Char 13")
	ht.set(20, "Char 20")
	ht.set(-9, "Char -9")
	fmt.Println(ht.getSize())
	fmt.Println(ht.get(0))
	fmt.Println(ht.get(13))
	fmt.Println(ht.get(20))
	fmt.Println(ht.get(-2))
	ht.remove(20)
	fmt.Println(ht.get(20))
	fmt.Println(ht.getSize())
}
