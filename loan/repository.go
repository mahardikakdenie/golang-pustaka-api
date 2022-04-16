package loan

import (
	"pustaka-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Loan, error)
	Create(loan entity.Loan) (entity.Loan, error)
	FindById(Id int) (entity.Loan, error)
	UpdateLoan(loan entity.Loan) (entity.Loan, error)
	Destroy(loan entity.Loan) (entity.Loan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Loan, error) {
	var loans []entity.Loan
	err := r.db.Preload("User").Preload("Book").Find(&loans).Error

	return loans, err
}

func (r *repository) Create(Loan entity.Loan) (entity.Loan, error) {
	err := r.db.Create(&Loan).Error

	return Loan, err
}

func (r *repository) FindById(Id int) (entity.Loan, error) {
	var Loan entity.Loan
	err := r.db.Preload("User").Preload("Book").Find(&Loan, Id).Error

	return Loan, err
}

func (r *repository) UpdateLoan(Loan entity.Loan) (entity.Loan, error) {
	err := r.db.Save(&Loan).Error

	return Loan, err
}

func (r *repository) Destroy(Loan entity.Loan) (entity.Loan, error) {
	err := r.db.Delete(&Loan).Error

	return Loan, err
}
