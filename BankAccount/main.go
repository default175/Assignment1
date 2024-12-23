package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Account Number: ")
	accountNumber, _ := reader.ReadString('\n')
	accountNumber = strings.TrimSpace(accountNumber)

	fmt.Print("Enter Holder Name: ")
	holderName, _ := reader.ReadString('\n')
	holderName = strings.TrimSpace(holderName)

	account := BankAccount{
		AccountNumber: accountNumber,
		HolderName:    holderName,
		Balance:       0.0,
	}

	for {
		fmt.Println("\n1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Process Transactions")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Print("Enter amount to deposit: ")
			amountInput, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)

			err := account.Deposit(amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Amount deposited successfully.")
			}

		case "2":
			fmt.Print("Enter amount to withdraw: ")
			amountInput, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)

			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Amount withdrawn successfully.")
			}

		case "3":
			fmt.Printf("Current Balance: %.2f Tenge\n", account.GetBalance())

		case "4":
			fmt.Print("Enter transactions (comma-separated, positive for deposits, negative for withdrawals): ")
			transactionsInput, _ := reader.ReadString('\n')
			transactionsStr := strings.Split(strings.TrimSpace(transactionsInput), ",")

			var transactions []float64
			for _, str := range transactionsStr {
				transaction, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
				if err == nil {
					transactions = append(transactions, transaction)
				}
			}

			Transaction(&account, transactions)
			fmt.Println("Transactions processed successfully.")

		case "5":
			fmt.Println("Exiting the Banking System.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
