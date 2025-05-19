package controllers

import (
    //"encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    //"example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"
    //"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
    Service services.BookServiceInterface
}

func NewBookController(service services.BookServiceInterface) *BookController {
    return &BookController{Service: service}
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    err := c.Service.DeleteBookByID(id)
    if err != nil {
        http.Error(w, "Error al eliminar libro", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}


func (c *BookController) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    err := c.Service.DeleteBookByID(id)
    if err != nil {
        http.Error(w, "Error al eliminar libro", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
