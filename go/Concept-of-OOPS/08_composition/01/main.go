package main

import "fmt"

type Engine struct {
	horsepower float64
}

func (e *Engine) Start() {
	fmt.Printf("Car has %0.2f HP\n",e.horsepower)
}

type Wheel struct {
	class string
}

func (w *Wheel) Rotate() {
	fmt.Printf("These wheels are %s type\n",w.class)
}

type Transmission struct{
	class string
}

func (t *Transmission) ShiftingGear() {
	fmt.Printf("Car is using %s type of transmission\n", t.class)
}

type Car struct {
	brandName string
	engine Engine
	wheel Wheel
	transmission Transmission
}

func (c *Car) Running() {
	fmt.Printf("The car name is %s",c.brandName)
	fmt.Println("\nThis car has following details -")
	fmt.Println("--------------------------------------")
	c.engine.Start()
	c.wheel.Rotate()
	c.transmission.ShiftingGear()
}

func main() {

	var engine = Engine{horsepower:172.89}

	var wheel = Wheel{class:"Alloy"}

	var transmission = Transmission{class:"Automatic"}

	car := Car{
		brandName: "Toyota",
		engine: engine,
		wheel: wheel,
		transmission: transmission,
	}

	car.Running()
	
}
