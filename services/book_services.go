package services

import (
    "fmt"
    //"example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)

type BookServiceInterface interface {
    DeleteBookByID(id string) error
}

type BookService struct {
    repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
    return &BookService{repo}
}

func (s *BookService) DeleteBookByID(id string) error {
    deleted, err := s.repo.RemoveBookByID(id)
    if err != nil {
        return err
    }
    if !deleted {
        return fmt.Errorf("libro no encontrado")
    }
    return nil
}