package use_case

import (
	"errors"

	"github.com/t-sakoda/expense-tracker/domain"
)

type IAddExpenseUseCase interface {
	Execute(description string, amount float64) (uint64, error)
}

type AddExpenseUseCase struct {
	Repo domain.ExpenseRepository
}

func (uc *AddExpenseUseCase) Execute(description string, amount float64) (uint64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid expense amount")
	}
	if description == "" {
		return 0, errors.New("description is required")
	}
	expense := domain.NewExpense(description, amount)
	err := uc.Repo.Create(expense)
	return expense.Id, err
}
