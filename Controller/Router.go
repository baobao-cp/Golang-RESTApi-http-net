package Controller

import "net/http"

type Router interface {
	HandleBookRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	bookController BookController
}

func NewRouter(bkc BookController) Router {

	return &router{
		bkc,
	}
}

func (rt *router) HandleBookRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rt.bookController.GetAllBooks(w, r)
	case "POST":
		rt.bookController.AddBook(w, r)
	case "PUT":
		rt.bookController.UpdateBookById(w, r)
	case "DELETE":
		rt.bookController.DeleteBookById(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}
