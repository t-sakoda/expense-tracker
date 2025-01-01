package use_case

import (
	"errors"

	"github.com/t-sakoda/expense-tracker/domain"
)

type AddExpenseUseCase struct {
	repo domain.ExpenseRepository
}

func NewAddExpenseUseCase(repo domain.ExpenseRepository) *AddExpenseUseCase {
	return &AddExpenseUseCase{repo: repo}
}

func (uc *AddExpenseUseCase) Execute(description string, amount float64) error {
	if amount <= 0 {
		return errors.New("invalid expense amount")
	}
	if description == "" {
		return errors.New("description is required")
	}
	expense := domain.NewExpense(description, amount)
	err := uc.repo.Create(expense)
	return err
}
