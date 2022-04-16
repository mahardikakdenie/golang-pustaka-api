package loan

import (
	"fmt"
	"pustaka-api/entity"
)

type Service interface {
	FindAll() ([]entity.Loan, error)
	Create(request LoanRequest) (entity.Loan, error)
	FindById(Id int) (entity.Loan, error)
	Update(request LoanRequest, id int) (entity.Loan, error)
	DestroyLoan(id int) (entity.Loan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.Loan, error) {
	return s.repository.FindAll()
}

func (s *service) Create(request LoanRequest) (entity.Loan, error) {
	loan, err := s.repository.Create(entity.Loan{
		UserId:     request.UserId,
		BookId:     request.BookId,
		Status:     request.Status,
		DateReturn: request.DateReturn,
	})

	return loan, err
}

func (s *service) FindById(Id int) (entity.Loan, error) {
	return s.repository.FindById(Id)
}

func (s *service) Update(request LoanRequest, id int) (entity.Loan, error) {
	dataLoanId, err := s.repository.FindById(id)
	if err != nil {
		fmt.Println(err)
	}

	dataLoanId.BookId = request.BookId
	dataLoanId.UserId = request.UserId
	dataLoanId.Status = request.Status
	dataLoanId.DateReturn = request.DateReturn

	updateDataloan, err := s.repository.UpdateLoan(dataLoanId)

	return updateDataloan, err

}

func (s *service) DestroyLoan(id int) (entity.Loan, error) {
	loan, err := s.repository.FindById(id)
	if err != nil {
	}
	deleteData, err := s.repository.Destroy(loan)

	return deleteData, err
}
