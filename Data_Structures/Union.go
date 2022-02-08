package main

import (
	"fmt"
)

type Unioner interface {
	add()
	find()
	union()
}

type UnionElement struct {
	element string
	group   *UnionElement
}

type Union map[string]*UnionElement

func (u Union) add(ele string) {
	u[ele] = &UnionElement{ele, nil}
}

func (u Union) find(ele string) string {
	if u[ele].group == nil {
		return u[ele].element
	} else {
		root := u.find(u[ele].group.element)
		u[ele].group = u[root]
		return root
	}
}

func (u Union) union(ele1, ele2 string) {
	if ele1 == ele2 {
		fmt.Println("Cannot union the same element.")
	} else {
		gp1, gp2 := u.find(ele1), u.find(ele2)
		if gp1 != gp2 {
			u[ele1].group = u[ele2]
		}
	}
}

func main() {
	fmt.Println("------Union Testing------")
	un := Union{}
	un.add("A01")
	un.add("A02")
	un.add("A03")
	un.add("B01")
	un.add("B02")
	un.add("B03")
	un.add("C01")
	un.add("C02")
	un.add("C03")
	un.union("A03", "A02")
	un.union("A02", "A01")
	un.union("C03", "C01")
	fmt.Printf("Current root for A01: %s, A02: %s, A03: %s, C03: %s, B01: %s\n", un.find("A01"), un.find("A02"), un.find("A03"), un.find("C03"), un.find("B01"))
}
