package main

import (
	"errors"
	"fmt"
)


type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}


func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than 0")
	}
	ba.Balance += amount
	return nil
}


func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than 0")
	}
	if ba.Balance < amount {
		return errors.New("insufficient funds")
	}
	ba.Balance -= amount
	return nil
}


func (ba *BankAccount) GetBalance() float64 {
	return ba.Balance
}


func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Println("Deposit Error:", err)
			}
		} else {
			err := account.Withdraw(-amount) // Отрицательные суммы - это снятие
			if err != nil {
				fmt.Println("Withdrawal Error:", err)
			}
		}
	}
}
