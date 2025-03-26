package main

import "fmt"

type Animal interface {
	speak()
}

type Dog struct {
	name string
}

func (d Dog) speak() {
	fmt.Println(d.name, "is barking")
}

type Cat struct {
	name string
}

func (c Cat) speak() {
	fmt.Println(c.name,"is mewing")
}

func main() {

	var a Animal = Dog{name:"Tommy",}

	a.speak()

	var a1 Animal = Cat{name:"Tom",}

	a1.speak()
	
}
