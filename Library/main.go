package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	library := NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose:")
		fmt.Println("1. Add")
		fmt.Println("2. Borrow")
		fmt.Println("3. Return")
		fmt.Println("4. List")
		fmt.Println("5. Exit")
		fmt.Print("Ваш выбор: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "Add":
			fmt.Print("Enter ID: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			fmt.Print("input name: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			fmt.Print("input author: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())
			book := Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			library.AddBook(book)

		case "Borrow":
			fmt.Print("input author: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			if err := library.BorrowBook(id); err != nil {
				fmt.Println("Ошибка:", err)
			}

		case "Return":
			fmt.Print("Return book- enter id: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			if err := library.ReturnBook(id); err != nil {
				fmt.Println("Err:", err)
			}

		case "List":
			library.ListBooks()

		case "Exit":
			fmt.Println("End programm.")
			return

		default:
			fmt.Println("Invalid choise.")
		}
	}
}
