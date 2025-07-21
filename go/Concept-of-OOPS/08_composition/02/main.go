package main

import "fmt"

type Engine interface {
	Start()
}

type Petrol struct {}

func (p Petrol) Start() {
	fmt.Println("Petrol car has been Started")
}

type Diesel struct {}

func (d Diesel) Start() {
	fmt.Println("Diesel car has been Started")
}

type Car struct {
	engine Engine
}

func (c Car) Running() {
	c.engine.Start()
	fmt.Println("Car is running")
}

func main() {
	p := Petrol{}

	d := Diesel{}

	c := Car{engine: p}

	c.Running()

	c1 := Car{engine: d}
	
	c1.Running()
}
