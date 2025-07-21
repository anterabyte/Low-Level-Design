package main

import "fmt"

type Vehicle interface {
	Start()
	DisplayBrand()
}


type Car struct {
	brand string
}

func (c Car) Start() {
	fmt.Println("Car is starting....")
}

func (c Car) DisplayBrand() {
	fmt.Println(c.brand, "is the brand name")
}


func main() {

	var name Vehicle = Car{brand:"Toyota",}

	name.Start()

	name.DisplayBrand()

}
