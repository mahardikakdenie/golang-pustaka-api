package book

import (
	entity "pustaka-api/entity"
)

type Service interface {
	FindAll(entity string) ([]entity.Book, error)
	FindById(id int) (entity.Book, error)
	Create(book BookRequest) (entity.Book, error)
	Update(book BookRequest, id int) (entity.Book, error)
	Destroy(id int) (entity.Book, error)
	FileUpload(id int, url string) (entity.Book, error)
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

func (s *service) Update(book BookRequest, id int) (entity.Book, error) {
	indexBook, _ := s.repository.FindById(id)

	indexBook.Title = book.Title
	indexBook.Price = book.Price
	indexBook.Substitle = book.Substitle
	indexBook.Rating = book.Rating
	indexBook.Description = book.Description
	indexBook.UserId = book.UserId

	updateBook, err := s.repository.Update(indexBook)

	return updateBook, err
}

func (s *service) Destroy(id int) (entity.Book, error) {
	books, _ := s.repository.FindById(id)
	newBook, err := s.repository.Destroy(books)

	return newBook, err

}

func (s *service) FileUpload(id int, url string) (entity.Book, error) {
	books, _ := s.repository.FindById(id)
	books.UrlImage = url

	newBook, err := s.repository.FileUpload(books)

	return newBook, err
}
