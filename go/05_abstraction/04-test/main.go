package main

import "fmt"

type paymentSystem interface {

	Payment(amount float64)
}

type PaybyPaypal struct {}

func (p PaybyPaypal) Payment(amount float64) {

	fmt.Printf("The amount paid is %.2f from paypal\n",amount)
	
}

type PayByPaytm struct {}

func (p PayByPaytm) Payment(amount float64) {
	fmt.Printf("The amount paid is %.2f from paytm\n",amount)
}

func main() {


	var sys paymentSystem = PaybyPaypal{}

	sys.Payment(156.43)

	var sys1 paymentSystem = PayByPaytm{}

	sys1.Payment(678.32)
}
