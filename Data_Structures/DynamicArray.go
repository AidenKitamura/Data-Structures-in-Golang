package main

import (
	"fmt"
	"strconv"
)

type DynamicArrayer interface {
	add()
	get()
	set()
	getSize()
	remove()
	isEmpty()
}

type DynamicArray struct {
	size  int
	array []int
}

func (x *DynamicArray) add(ele int) {
	if x.size == len(x.array) {
		newArray := make([]int, 2*len(x.array))
		copy(newArray, x.array)
		newArray[len(x.array)] = ele

		x.array = newArray
		x.size += 1
		fmt.Printf("New element added: %d, current size: %d\n", ele, x.size)
	} else {
		x.array[x.size] = ele
		x.size += 1
		fmt.Printf("New element added: %d, current size: %d\n", ele, x.size)
	}
}

func (x *DynamicArray) get(index int) int {
	if index >= x.size || index < 0 {
		fmt.Println("Error, index out of range")
		return 0
	} else {
		return x.array[index]
	}
}

func (x *DynamicArray) set(index, ele int) {
	if index >= x.size || index < 0 {
		fmt.Println("Error, index out of range")
	} else {
		x.array[index] = ele
	}
	fmt.Printf("New element set: %d at position %d\n", ele, index)
}

func (x *DynamicArray) getSize() int {
	return x.size
}

func (x *DynamicArray) remove(index int) {
	if index >= x.size || index < 0 {
		fmt.Println("Error, index out of range")
	} else {
		for i := index; i < x.size-1; i += 1 {
			x.array[i] = x.array[i+1]
		}
		x.size -= 1
		fmt.Printf("Element at Position %d deleted\n", index)
	}
}

func (x *DynamicArray) isEmpty() bool {
	if x.size == 0 {
		return true
	} else {
		return false
	}
}

func (x *DynamicArray) String() string {
	info := fmt.Sprintf("Size of the Dynamic Array:%d, current elements: ", x.size)
	if !x.isEmpty() {
		for i := 0; i < x.size; i += 1 {
			info += strconv.Itoa(x.array[i]) + ", "
		}
	}
	return info
}

func main() {
	nArray := DynamicArray{}
	nArray.array = make([]int, 4)
	nArray.add(2)
	nArray.add(3)
	fmt.Printf(nArray.String())
	nArray.add(1)
	nArray.add(765)
	nArray.add(6788)
	nArray.add(798)
	fmt.Println(strconv.Itoa(nArray.get(3)))
	fmt.Printf("Current Array Size: %d\n", nArray.getSize())
	nArray.remove(2)
	fmt.Printf("Current Array Size: %d\n", nArray.getSize())
	nArray.remove(9)
	fmt.Printf(nArray.String())
}
