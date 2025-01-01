package infra_test

import "github.com/t-sakoda/expense-tracker/domain"

type MockExpenseRepository struct {
	insertCallCount int
}

func (r *MockExpenseRepository) InsertCallCount() int {
	return r.insertCallCount
}

func (r *MockExpenseRepository) GenerateNewId() (uint64, error) {
	return 1, nil
}

func (r *MockExpenseRepository) Save(expense *domain.Expense) error {
	r.insertCallCount++
	return nil
}
