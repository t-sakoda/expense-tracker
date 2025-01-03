package infra

import (
	"github.com/t-sakoda/expense-tracker/domain"
)

type MockExpenseRepository struct {
	SaveFunc               func(expense *domain.Expense) error
	GenerateNewIdFunc      func() (uint64, error)
	SaveCallCount          int
	GenerateNewIdCallCount int
}

func (m *MockExpenseRepository) Save(expense *domain.Expense) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(expense)
	}
	// default implementation
	m.SaveCallCount++
	return nil
}

func (m *MockExpenseRepository) GenerateNewId() (uint64, error) {
	if m.GenerateNewIdFunc != nil {
		return m.GenerateNewIdFunc()
	}
	// default implementation
	m.GenerateNewIdCallCount++
	return 1, nil
}
