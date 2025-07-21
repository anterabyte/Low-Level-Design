package main

import "fmt"

type Account struct {
	name string
	balance float64
}

// Constructor Function
func NewAccount(name string, balance float64) *Account {
	return &Account{name: name, balance: balance}
}

// Getter
func (a *Account) GetBalance() string{
	return fmt.Sprintln("Your current balance is",a.balance)
}

// Setter
func (a *Account) SetBalance(amount float64) {
	if amount > 0 {
		a.balance += amount  // Add Money to the Bank
	}else {
		fmt.Println("theres is nothing to add in the bank account")
	} 
 }

func main(){
	acc := NewAccount("Ank", 1122.90)

	fmt.Println (acc.GetBalance())

	acc.SetBalance(567.89)

	fmt.Println(acc.GetBalance())
}
