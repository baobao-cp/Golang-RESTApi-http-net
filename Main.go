package main

import (
	"GolangTrain/Controller"
	"GolangTrain/Model/Repository"
	"github.com/fatih/color"
	"net/http"
	"time"
)

var (
	bookRepository = Repository.NewBookRepository()
	bookController = Controller.NewBookController(bookRepository)
	router         = Controller.NewRouter(bookController)
)

func main() {

	server := http.NewServeMux()
	server.HandleFunc("/book/", router.HandleBookRequest)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
	color.CyanString("Server is running on port 8080 : " + time.Now().String())
}
