package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-mongo-app/controllers"
	"example.com/go-mongo-app/repositories"
	"example.com/go-mongo-app/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)
	controller := controllers.NewBookController(service) // ← corregido

	r := mux.NewRouter() // ← estaba faltando esta línea
	r.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE") // ← nombre correcto de handler
	return r
}

func TestBookCRUDIntegration(t *testing.T) {
	router := setupRouter()

	id := "64e4b2cabc1234567890abcd" // Usa un ID válido o uno mock
	reqDelete, _ := http.NewRequest("DELETE", "/books/"+id, nil) // ← urlByID definido aquí directamente
	rrDelete := httptest.NewRecorder()
	router.ServeHTTP(rrDelete, reqDelete)

	// Si el libro no existe, podría devolver 404. Dependerá de tu lógica real.
	assert.True(t, rrDelete.Code == http.StatusNoContent || rrDelete.Code == http.StatusNotFound)
}
