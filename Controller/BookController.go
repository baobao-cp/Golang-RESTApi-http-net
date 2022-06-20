package Controller

import (
	"GolangTrain/Controller/DTO"
	"GolangTrain/Model/Entiry"
	"GolangTrain/Model/Repository"
	"GolangTrain/Validate"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BookController interface {
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	AddListBook(w http.ResponseWriter, r *http.Request)
	DeleteBookById(w http.ResponseWriter, r *http.Request)
	UpdateBookById(w http.ResponseWriter, r *http.Request)
}

type bookController struct {
	bookRepository Repository.BookRepository
}

func NewBookController(bookRepository Repository.BookRepository) BookController {
	return &bookController{bookRepository}
}

func (bk *bookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := bk.bookRepository.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// To JSON
	responseBook, _ := json.MarshalIndent(books, "", " ")

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBook)
	color.Blue("GetAllBooks request successfully  " + time.Now().String())

}

func (bk *bookController) AddBook(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var book Entiry.Book
	err := json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON syntax error in request body"))
	}
	fmt.Println(book)

	if errors, haveError := Validate.CheckBook(book); haveError {
		indent, _ := json.MarshalIndent(errors, "", " ")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(indent)
		return
	}

	book, err = bk.bookRepository.AddBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// To JSON
	var dataResponse = DTO.DataResponseDTO{
		Message: "book success added",
		Data:    book,
	}
	responseBook, _ := json.MarshalIndent(dataResponse, "", " ")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBook)
}

func (bk *bookController) AddListBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		var responseDTO = DTO.DataResponseDTO{
			Message: "Error while reading request body",
			Data:    nil,
		}
		errorResponse, err := json.MarshalIndent(responseDTO, "", " ")
		if err != nil {
			return
		}
		color.Red("Error reading body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResponse)
		return
	}
	var listBooks []Entiry.Book
	err = json.Unmarshal(body, &listBooks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON syntax error in request body"))
		return
	}
	listBooks, err = bk.bookRepository.AddListBook(listBooks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// To JSON
	var dataResponse = DTO.DataResponseDTO{
		Message: "books success added",
		Data:    listBooks,
	}
	responseBooks, _ := json.MarshalIndent(dataResponse, "", " ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseBooks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	color.Blue("AddListBook request successfully  " + time.Now().String())
}

func (bk *bookController) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	var id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/book/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
	}
	err, existId := bk.bookRepository.DeleteBookById(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !existId {
		w.Write([]byte("Book not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted"))

}

func (bk *bookController) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var path = strings.TrimPrefix(r.URL.Path, "/book/")
	fmt.Println(path)
	var id, err = strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var book Entiry.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON syntax error in request body"))
	}
	book, err, existId := bk.bookRepository.UpdateBookById(int64(id), book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !existId {
		w.Write([]byte("Book not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// To JSON
	var dataResponse = DTO.DataResponseDTO{
		Message: "book success updated",
		Data:    book,
	}
	responseBook, _ := json.MarshalIndent(dataResponse, "", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBook)
}
