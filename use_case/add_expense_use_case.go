package use_case

import (
	"errors"
	"fmt"

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
	id, err := uc.Repo.GenerateNewId()
	if err != nil {
		return 0, fmt.Errorf("failed to generate new ID: %w", err)
	}
	expense := &domain.Expense{
		Id:          id,
		Description: description,
		Amount:      amount,
	}
	errSave := uc.Repo.Save(expense)
	if errSave != nil {
		return 0, fmt.Errorf("failed to save expense: %w", errSave)
	}
	return id, nil
}
