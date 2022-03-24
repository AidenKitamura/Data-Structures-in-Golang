package util

import (
	"fmt"
)

type BinarySearchTreeer interface {
	insert()
	remove()
	find()
	getSize()
	isEmpty()
	traverse()
}

type Node struct {
	element    int
	leftChild  *Node
	rightChild *Node
	parent     *Node
}

type BinarySearchTree struct {
	size int
	root *Node
}

func (b *BinarySearchTree) insert(ele int) {
	if b.isEmpty() {
		b.root = &Node{ele, nil, nil, nil}
		b.size += 1
	} else {
		b.root._insert(ele)
		b.size += 1
	}
}

// Helper function
func (n *Node) _insert(ele int) {
	if ele <= n.element {
		if n.leftChild == nil {
			n.leftChild = &Node{ele, nil, nil, n}
		} else {
			n.leftChild._insert(ele)
		}
	} else {
		if n.rightChild == nil {
			n.rightChild = &Node{ele, nil, nil, n}
		} else {
			n.rightChild._insert(ele)
		}
	}
}

func (b *BinarySearchTree) remove(ele int) {
	status, addr := b.find(ele)
	if status == false {
		fmt.Println("Error, element not in the tree, cannot remove")
	} else {
		if addr.leftChild == nil && addr.rightChild == nil {
			if addr.parent.leftChild == addr {
				addr.parent.leftChild = nil
			} else {
				addr.parent.rightChild = nil
			}
		} else if addr.leftChild == nil {
			if addr.parent.leftChild == addr {
				addr.parent.leftChild = addr.rightChild
				addr.rightChild.parent = addr.parent
			} else {
				addr.parent.rightChild = addr.rightChild
				addr.rightChild.parent = addr.parent
			}
		} else if addr.rightChild == nil {
			if addr.parent.leftChild == addr {
				addr.parent.leftChild = addr.leftChild
				addr.leftChild.parent = addr.parent
			} else {
				addr.parent.rightChild = addr.leftChild
				addr.leftChild.parent = addr.parent
			}
		} else {
			sub := addr.leftChild._rightMostChild()
			addr.element = sub.element
			sub.remove()
		}
		b.size -= 1
	}
}

// Helper Function
func (n *Node) remove() {
	if n.leftChild == nil && n.rightChild == nil {
		if n.parent.leftChild == n {
			n.parent.leftChild = nil
		} else {
			n.parent.rightChild = nil
		}
	} else if n.leftChild == nil {
		if n.parent.leftChild == n {
			n.parent.leftChild = n.rightChild
			n.rightChild.parent = n.parent
		} else {
			n.parent.rightChild = n.rightChild
			n.rightChild.parent = n.parent
		}
	} else if n.rightChild == nil {
		if n.parent.leftChild == n {
			n.parent.leftChild = n.leftChild
			n.leftChild.parent = n.parent
		} else {
			n.parent.rightChild = n.leftChild
			n.leftChild.parent = n.parent
		}
	}
}

func (n *Node) _leftMostChild() *Node {
	if n.leftChild == nil {
		return n
	} else {
		return n.leftChild._leftMostChild()
	}
}

func (n *Node) _rightMostChild() *Node {
	if n.rightChild == nil {
		return n
	} else {
		return n.rightChild._rightMostChild()
	}
}

func (b BinarySearchTree) find(ele int) (bool, *Node) {
	if b.isEmpty() {
		return false, nil
	} else {
		return b.root._find(ele)
	}
}

// Helper function
func (n *Node) _find(ele int) (bool, *Node) {
	if n == nil {
		return false, nil
	} else if n.element < ele {
		return n.rightChild._find(ele)
	} else if n.element > ele {
		return n.leftChild._find(ele)
	} else {
		return true, n
	}
}

func (b BinarySearchTree) getSize() int {
	return b.size
}

func (b BinarySearchTree) isEmpty() bool {
	return b.size == 0
}

/*
func main() {
	fmt.Println("------Binary Search Tree Testing------")
	bst := BinarySearchTree{}
	bst.insert(0)
	bst.insert(10)
	bst.insert(3)
	bst.insert(-5)
	fmt.Println(bst.getSize())
	fmt.Println(bst.root.element)
	fmt.Println(bst.root.leftChild.element)
	fmt.Println(bst.root.rightChild.element)
	fmt.Println(bst.root.rightChild.leftChild.element)
	_, addr := bst.find(3)
	fmt.Println(addr)
	bst.insert(1)
	bst.insert(-5)
	bst.insert(0)
	bst.insert(9)
	fmt.Println(bst.root.leftChild.rightChild)
	fmt.Println(bst.root.leftChild.leftChild)
	fmt.Println(bst.root.rightChild.leftChild.rightChild)
	bst.remove(0)
	fmt.Println(bst.getSize())
	bst.remove(0)
	fmt.Println(bst.getSize())
	fmt.Println(bst.root)
	status, _ := bst.find(0)
	fmt.Println(status)
	bst.remove(1000)
}
*/
