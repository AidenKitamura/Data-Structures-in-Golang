package util

import (
	"fmt"
)

type FenwickTreeer interface {
	construct()
	query()
	getSize()
	isEmpty()
}

type FenwickTree struct {
	size int
	tree []int
}

func (f *FenwickTree) construct(arr []int) {
	newArray := make([]int, len(arr))
	copy(newArray, arr)
	n := len(arr)
	for i := 1; i <= n; i++ {
		parent := _LSB(i)
		if parent+i <= n {
			newArray[parent+i-1] += newArray[i-1]
		}
	}
	f.size = n
	f.tree = newArray
}

func (f FenwickTree) query(i int) int {
	if i == 0 {
		return 0
	} else {
		return f.tree[i-1] + f.query(i-_LSB(i))
	}
}

func (f FenwickTree) getSize() int {
	return f.size
}

func (f FenwickTree) isEmpty() bool {
	return f.size == 0
}

// Testing Function
func (f FenwickTree) print() {
	for i := 0; i < f.getSize(); i++ {
		fmt.Println(f.tree[i])
	}
}

func _LSB(e int) int {
	res := 0
	for num := e; num != 0; num /= 2 {
		if num%2 == 0 {
			res += 1
		} else {
			break
		}
	}
	return 1 << res
}

func main() {
	fmt.Println("------Fenwick Tree Testing------")
	f := FenwickTree{}
	f.construct([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.print()
	fmt.Println(f.query(10))
	fmt.Println(f.query(9))
	fmt.Println(f.query(8))
}
