package main

import "fmt"

type Book struct {
	Title       string
	Author      string
	IsAvailable bool
}

func (book *Book) DisplayInfo() {

	fmt.Printf("\nBook Title: %s \nAuthor : %s \nAvailable : %s\n\n",
		book.Title, book.Author,
		map[bool]string{true: "Yes", false: "No"}[book.IsAvailable])
}
func (book *Book) UpdateTitle(title string) {
	book.Title = title
	fmt.Println("Updated Title")
}
func (book *Book) UpdateAuthor(author string) {
	book.Author = author
	fmt.Println("Updated Author")
}

func (book *Book) BorrowBook() {
	if book.IsAvailable {
		book.IsAvailable = false
		fmt.Println("Book has been borrowed.")
	} else {
		fmt.Println("This book is already borrowed.")
	}
}
func (book *Book) ReturnBook() {
	if book.IsAvailable {
		fmt.Println("This book is already available.")
	} else {
		book.IsAvailable = true
		fmt.Println("Book has been returned.")
	}
}

func main() {
	myBook := Book{Title: "The Stranger", Author: "Albert Camus", IsAvailable: true}

	myBook.DisplayInfo()
	myBook.UpdateTitle("The Fall")
	myBook.DisplayInfo()
	myBook.UpdateAuthor("Albert Camus Jr.")
	myBook.DisplayInfo()
	myBook.BorrowBook()
	myBook.DisplayInfo()
	myBook.BorrowBook()
	myBook.DisplayInfo()
	myBook.ReturnBook()
	myBook.DisplayInfo()
	myBook.ReturnBook()
	myBook.DisplayInfo()

}
