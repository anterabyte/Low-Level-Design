package main

import "fmt"

type Car struct {
	// Attributes
	Color string
	Make  string
	Model string
	Year  string
}

// Method to display car details
func (c Car) DisplayInfo() {
	fmt.Println("Car Make: " + c.Make)
	fmt.Println("Car Model: " + c.Model)
	fmt.Println("Car Year: " + c.Year)
	fmt.Println("Car Color: " + c.Color)
}
