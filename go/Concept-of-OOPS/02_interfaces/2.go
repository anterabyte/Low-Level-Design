package main

import "fmt"

type fly interface {

	flyingcars()
}

type drive interface {

	drivingcars()
	
}

type cars struct {}

func (c cars) flyingcars(){

	fmt.Println("I am flying in a car")
}

func (c cars) drivingcars(){

	fmt.Println("I am driving a car")
}

func main(){

	var f fly = cars{}

	f.flyingcars()

	var d drive = cars{}

	d.drivingcars()

}
