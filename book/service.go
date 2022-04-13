package book

import (
	"encoding/json"
	entity "pustaka-api/entity"
)

type Service interface {
	FindAll(entity string) ([]entity.Book, error)
	FindById(id int) (entity.Book, error)
	Create(book BookRequest) (entity.Book, error)
	Update(id int, book BookRequest) (entity.Book, error)
	Destroy(id int) (entity.Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(entity string) ([]entity.Book, error) {
	return s.repository.FindAll(entity)
}

func (s *service) FindById(Id int) (entity.Book, error) {
	return s.repository.FindById(Id)
}

func (s *service) Create(book BookRequest) (entity.Book, error) {
	books, err := s.repository.Create(entity.Book{
		Title:       book.Title,
		Price:       book.Price,
		Substitle:   book.Substitle,
		Rating:      book.Rating,
		Description: book.Description,
		UserId:      book.UserId,
	})

	return books, err
}

func (s *service) Update(id int, book BookRequest) (entity.Book, error) {
	indexBook, _ := s.repository.FindById(id)

	titleBook := ternaryTitle(book, indexBook)
	priceBook := ternaryPrice(book, indexBook)
	substitleBook := ternarySubstitle(book, indexBook)
	ratingBook := ternaryRating(book, indexBook)
	descriptionBook := ternaryDescription(book, indexBook)

	indexBook.Title = titleBook
	indexBook.Price = priceBook
	indexBook.Substitle = substitleBook
	indexBook.Rating = ratingBook
	indexBook.Description = descriptionBook

	updateBook, err := s.repository.Update(indexBook)

	return updateBook, err
}

func (s *service) Destroy(id int) (entity.Book, error) {
	books, _ := s.repository.FindById(id)
	newBook, err := s.repository.Destroy(books)

	return newBook, err

}

func ternaryTitle(request BookRequest, indexBook entity.Book) string {
	if request.Title != "" {
		indexBook.Title = request.Title
	}
	return indexBook.Title
}

func ternaryPrice(request BookRequest, indexBook entity.Book) json.Number {
	if request.Price != "" {
		indexBook.Price = request.Price
	}
	return indexBook.Price
}

func ternarySubstitle(request BookRequest, indexBook entity.Book) string {
	if request.Substitle != "" {
		indexBook.Substitle = request.Substitle
	}
	return indexBook.Substitle
}

func ternaryRating(request BookRequest, indexBook entity.Book) json.Number {
	if request.Rating != "" {
		indexBook.Rating = request.Rating
	}
	return indexBook.Rating
}

func ternaryDescription(request BookRequest, indexBook entity.Book) string {
	if request.Description != "" {
		indexBook.Description = request.Description
	}
	return indexBook.Description
}
