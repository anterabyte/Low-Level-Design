package main

import "fmt"

type name struct {

	first string
	last string
}

func (n name) eat() {
	fmt.Println(n.first, n.last,"was eating some food over there" )
}

type human struct {
	name
}

func (h human) flies() {
	fmt.Println(h.first, h.last,"was flying")
}

func main() {

	myhuman := human{name{first:"Kal",last:"kothri",}}

	myhuman.eat()
	myhuman.flies()
}
