package main

import (
	"errors"
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]Book)}
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
	fmt.Printf("Book '%s' Added library.\n", book.Title)
}

func (l *Library) BorrowBook(id string) error {
	book, exists := l.Books[id]
	if !exists {
		return errors.New("book with this ID not found")
	}
	if book.IsBorrowed {
		return errors.New("Book already borrowed")
	}
	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Printf("Book '%s' borrowed.\n", book.Title)
	return nil
}

func (l *Library) ReturnBook(id string) error {
	book, exists := l.Books[id]
	if !exists {
		return errors.New("Book with this ID not found")
	}
	if !book.IsBorrowed {
		return errors.New("Book already borrowed")
	}
	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Printf("Book '%s' returned.\n", book.Title)
	return nil
}

func (l *Library) ListBooks() {
	fmt.Println("List of available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("- ID: %s, Name: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
