package Repository

import (
	"GolangTrain/Model/Entiry"
	"fmt"
	"github.com/fatih/color"
)

type BookRepository interface {
	GetAllBooks() ([]Entiry.Book, error)
	AddBook(b Entiry.Book) (Entiry.Book, error)
	AddListBook(b []Entiry.Book) ([]Entiry.Book, error)
	DeleteBookById(id int64) (error, bool)
	UpdateBookById(id int64, b Entiry.Book) (Entiry.Book, error, bool)
}

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (bk *BookRepositoryImpl) GetAllBooks() ([]Entiry.Book, error) {
	var books []Entiry.Book

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		color.Red("Error while fetching books")
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var book Entiry.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			color.Red("Error while fetching books")
			panic(err.Error())
		}
		books = append(books, book)
	}
	return books, nil
}

func (bk *BookRepositoryImpl) AddBook(b Entiry.Book) (Entiry.Book, error) {
	var query = "INSERT INTO `books` (title, author) VALUES (?, ?)"
	insert, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	result, err := insert.Exec(b.Title, b.Author)
	if err != nil {
		fmt.Print(err.Error())
		return Entiry.Book{}, err
	}
	id, _ := result.LastInsertId()

	var responseBook = Entiry.Book{
		Id:     int(id),
		Title:  b.Title,
		Author: b.Author,
	}
	color.Blue("Book added successfully")
	return responseBook, nil

}

func (bk *BookRepositoryImpl) AddListBook(b []Entiry.Book) ([]Entiry.Book, error) {
	var books []Entiry.Book
	var query = "INSERT INTO `books` (title, author) VALUES (?, ?)"
	insert, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	for _, v := range b {
		result, err := insert.Exec(v.Title, v.Author)
		if err != nil {
			fmt.Print(err.Error())
			return nil, err
		}
		responseId, _ := result.LastInsertId()
		books = append(books, Entiry.Book{
			Id:     int(responseId),
			Title:  v.Title,
			Author: v.Author,
		})
	}
	color.Blue("Books added successfully")
	return books, nil
}

func (bk *BookRepositoryImpl) DeleteBookById(id int64) (error, bool) {
	var existQuery = "SELECT Id FROM `books` WHERE id = ?"
	existStmt, err := db.Query(existQuery, id)
	if err != nil {
		panic(err.Error())
	}
	defer existStmt.Close()
	var existId int64
	for existStmt.Next() {
		err := existStmt.Scan(&existId)
		if err != nil {
			panic(err.Error())
		}
	}
	if existId == 0 {
		color.Red("Book not found")
		return nil, false
	}

	var query = "DELETE FROM `books` WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Print(err.Error())
		return err, false
	}
	color.Blue("Book deleted successfully")
	return nil, true
}

func (bk *BookRepositoryImpl) UpdateBookById(id int64, b Entiry.Book) (Entiry.Book, error, bool) {
	var existQuery = "SELECT Id FROM `books` WHERE id = ?"
	existStmt, err := db.Query(existQuery, id)
	if err != nil {
		panic(err.Error())
	}
	defer existStmt.Close()
	var existId int64
	for existStmt.Next() {
		err := existStmt.Scan(&existId)
		if err != nil {
			panic(err.Error())
		}
	}
	if existId == 0 {
		color.Red("Book not found")
		return Entiry.Book{}, nil, false
	}

	var query = "UPDATE `books` SET title = ?, author = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(b.Title, b.Author, id)
	if err != nil {
		fmt.Print(err.Error())
		return Entiry.Book{}, err, true
	}

	var responseBook = Entiry.Book{
		Id:     int(id),
		Title:  b.Title,
		Author: b.Author,
	}
	color.Blue("Book updated successfully")

	return responseBook, nil, true
}
