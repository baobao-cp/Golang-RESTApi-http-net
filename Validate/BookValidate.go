package Validate

import (
	"GolangTrain/Model/Entiry"
	"fmt"
)

func CheckBook(book Entiry.Book) (map[string]string, bool) {
	errors := make(map[string]string)
	if book.Title == "" {
		errors["title"] = "Title is required"
	}
	if book.Author == "" {
		errors["author"] = "Author is required"
	}
	fmt.Println(errors)
	return errors, len(errors) != 0
}
