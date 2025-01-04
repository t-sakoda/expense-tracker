package service

import (
	"errors"

	"github.com/t-sakoda/expense-tracker/domain"
)

var ErrInvalidParameter = errors.New("invalid parameter")
var ErrFailedToGenerateId = errors.New("failed to generate id")
var ErrFailedToSaveExpense = errors.New("failed to save expense")
var ErrExpenseNotFound = errors.New("expense not found")
var ErrFailedToDeleteExpense = errors.New("failed to delete expense")
var ErrFailedToListExpenses = errors.New("failed to list expenses")

type ExpenseServiceInterface interface {
	Add(description string, amount float64) (uint64, error)
	Update(id uint64, description string, amount float64) error
	Delete(id uint64) error
	List() ([]domain.Expense, error)
}

type ExpenseService struct {
	repo  domain.ExpenseRepository
	clock domain.Clock
}

func NewExpenseService(repo domain.ExpenseRepository, clock domain.Clock) *ExpenseService {
	return &ExpenseService{
		repo:  repo,
		clock: clock,
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
		Date:        s.clock.Now(),
	}
	if err := s.repo.Save(expense); err != nil {
		return 0, ErrFailedToSaveExpense
	}
	return id, nil
}

func (s *ExpenseService) Update(id uint64, description string, amount float64) error {
	if description == "" || amount <= 0 {
		return ErrInvalidParameter
	}

	expense, err := s.repo.FindById(id)
	if err != nil {
		return ErrExpenseNotFound
	}

	expense.Description = description
	expense.Amount = amount

	if err := s.repo.Save(expense); err != nil {
		return ErrFailedToSaveExpense
	}

	return nil
}

func (s *ExpenseService) Delete(id uint64) error {
	if _, err := s.repo.FindById(id); err != nil {
		return ErrExpenseNotFound
	}
	if err := s.repo.Delete(id); err != nil {
		return ErrFailedToDeleteExpense
	}

	return nil
}

func (s *ExpenseService) List() ([]domain.Expense, error) {
	expenses, err := s.repo.FindAll()
	if err != nil {
		return nil, ErrFailedToListExpenses
	}
	return expenses, nil
}
