package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"example.com/go-mongo-app/controllers" // ← IMPORTANTE: para usar NewBookController
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) DeleteBookByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestDeleteBook_Success(t *testing.T) {
	mockService := new(MockBookService)
	controller := controllers.NewBookController(mockService) // ← corregido

	id := primitive.NewObjectID().Hex()
	mockService.On("DeleteBookByID", id).Return(nil)

	req := httptest.NewRequest("DELETE", "/books/"+id, nil)
	req = mux.SetURLVars(req, map[string]string{"id": id})
	resp := httptest.NewRecorder()

	controller.DeleteBook(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteBook_Error(t *testing.T) {
	mockService := new(MockBookService)
	controller := controllers.NewBookController(mockService)

	id := primitive.NewObjectID().Hex()
	mockService.On("DeleteBookByID", id).Return(assert.AnError)

	req := httptest.NewRequest("DELETE", "/books/"+id, nil)
	req = mux.SetURLVars(req, map[string]string{"id": id})
	resp := httptest.NewRecorder()

	controller.DeleteBook(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}
