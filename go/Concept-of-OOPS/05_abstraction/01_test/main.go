package main

import "fmt"

type Vehicle interface {
	Car()
	Truck()
}

type Sound struct{}

func (s Sound) Car() {
	fmt.Println("Vroom, Vroom!")
}

func (s Sound) Truck() {
	fmt.Println("Honk, Honk!")
}

func main() {


	var s Sound

	s.Car()

	s.Truck()
	
}
