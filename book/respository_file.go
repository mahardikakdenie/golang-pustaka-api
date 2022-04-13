package book

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{db}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("Find All")

	return books, errors.New("asu")
}

func (r *fileRepository) FindById(Id int) (Book, error) {
	var book Book

	fmt.Println("Find By Id")

	return book, errors.New("asu")
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create")

	return book, errors.New("asu")
}
