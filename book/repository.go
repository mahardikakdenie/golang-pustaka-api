package book

import (
	entity "pustaka-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(entity string) ([]entity.Book, error)
	FindById(id int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Destroy(book entity.Book) (entity.Book, error)
	FileUpload(book entity.Book) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(entities string) ([]entity.Book, error) {
	var books []entity.Book
	var err error
	if entities == "" {
		err = r.db.Find(&books).Error
	} else {
		err = r.db.Preload("User").Find(&books).Error
	}
	// err := r.db.Raw("SELECT users.id as user_id, books.id FROM books INNER JOIN users ON books.user_id = users.id").Scan(&books).Error

	return books, err
}

func (r *repository) FindById(Id int) (entity.Book, error) {
	var books entity.Book
	err := r.db.Preload("User").Find(&books, Id).Error

	return books, err
}

func (r *repository) Create(books entity.Book) (entity.Book, error) {
	err := r.db.Create(&books).Error

	return books, err
}

func (r *repository) Update(books entity.Book) (entity.Book, error) {
	err := r.db.Save(&books).Error

	return books, err
}

func (r *repository) Destroy(books entity.Book) (entity.Book, error) {
	err := r.db.Delete(&books).Error

	return books, err
}

func (r *repository) FileUpload(books entity.Book) (entity.Book, error) {
	err := r.db.Save(&books).Error

	return books, err
}
