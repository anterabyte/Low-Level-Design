package main

import "fmt"

type vehical interface {
	start()
	stop()
}

type brand struct {
	car string
}

func (b brand) start() {

	fmt.Println("Start this car named -",b.car)

}

func (b brand) stop() {

	fmt.Println("Stop this car named -",b.car)

}

func main() {

	var v vehical = brand{car: "maruti",}

	v.start()
	v.stop()
}
 
