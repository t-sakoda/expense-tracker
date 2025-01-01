package infra_test

import "github.com/t-sakoda/expense-tracker/domain"

type MockExpenseRepository struct {
	insertCallCount int
}

func (r *MockExpenseRepository) InsertCallCount() int {
	return r.insertCallCount
}

func (r *MockExpenseRepository) Insert(expense *domain.Expense) error {
	r.insertCallCount++
	return nil
}
