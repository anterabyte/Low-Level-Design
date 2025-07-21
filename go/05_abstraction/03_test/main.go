package main

import "fmt"


type Animal interface {
	makeSound()
}

type Cat struct {}

func (c Cat) makeSound() {
	fmt.Println("It meows")
}

type Dog struct {}

func (d Dog) makeSound() {
	fmt.Println("It Barks")
}


func main() {

	myAnimal := Cat{}

	myAnimal.makeSound()

	myAnimal1 := Dog{}

	myAnimal1.makeSound()
	
}
