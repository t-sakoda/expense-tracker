package service

import (
	"errors"

	"github.com/t-sakoda/expense-tracker/domain"
)

var ErrInvalidParameter = errors.New("invalid parameter")
var ErrFailedToGenerateId = errors.New("failed to generate id")
var ErrFailedToSaveExpense = errors.New("failed to save expense")

type ExpenseServiceInterface interface {
	Add(description string, amount float64) (uint64, error)
	Update(id uint64, description string, amount float64) error
}

type ExpenseService struct {
	repo domain.ExpenseRepository
}

func NewExpenseService(repo domain.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) Add(description string, amount float64) (uint64, error) {
	if description == "" || amount <= 0 {
		return 0, ErrInvalidParameter
	}

	id, generateErr := s.repo.GenerateNewId()
	if generateErr != nil {
		return 0, ErrFailedToGenerateId
	}
	expense := &domain.Expense{
		Id:          id,
		Description: description,
		Amount:      amount,
	}
	if err := s.repo.Save(expense); err != nil {
		return 0, ErrFailedToSaveExpense
	}
	return id, nil
}

func (s *ExpenseService) Update(id uint64, description string, amount float64) error {
	return errors.New("not implemented")
}
