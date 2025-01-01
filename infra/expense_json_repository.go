package infra

import "github.com/t-sakoda/expense-tracker/domain"

type ExpenseJsonRepository struct {
}

func (repo ExpenseJsonRepository) GenerateNewId() uint64 {
	// TODO implement
	return 0
}

func (repo ExpenseJsonRepository) Insert(expense *domain.Expense) error {
	// TODO implement
	return nil
}
