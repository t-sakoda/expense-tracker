package infra

import "github.com/t-sakoda/expense-tracker/domain"

type ExpenseJsonRepository struct {
}

func (repo ExpenseJsonRepository) Create(expense *domain.Expense) error {
	panic("implement me")
}
