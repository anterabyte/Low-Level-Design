package main

import "fmt"

type BankAccount struct {
	name    string
	balance float64
}

// constructor function to intialize the BankAccount
func NewBankBalance(name string, amount float64) *BankAccount {
	return &BankAccount{name: name, balance: amount}
}

// Getter Method to access the Balance
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

// Setter Method to modify the Balance
func (b *BankAccount) DepositToAccount(amount float64) {
	if amount > 0 {
		b.balance += amount
		fmt.Printf("%.2f deposited to accout\n", amount)
	} else {
		fmt.Println("Invalid amount deposited")
	}
}

func main() {

	newAccount := NewBankBalance("Alice", 1100.23)

	fmt.Printf("Your current account balance is %0.2f\n", newAccount.balance)

	newAccount.DepositToAccount(4500.32)

	fmt.Printf("Your current account balance is %0.2f\n", newAccount.balance)

}
