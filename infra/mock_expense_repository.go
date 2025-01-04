package infra

import (
	"github.com/t-sakoda/expense-tracker/domain"
)

type MockExpenseRepository struct {
	SaveFunc          func(expense *domain.Expense) error
	GenerateNewIdFunc func() (uint64, error)
	FindByIdFunc      func(id uint64) (*domain.Expense, error)

	SaveCallCount          int
	GenerateNewIdCallCount int
	FindByIdCallCount      int
}

func (m *MockExpenseRepository) Save(expense *domain.Expense) error {
	m.SaveCallCount++
	if m.SaveFunc != nil {
		return m.SaveFunc(expense)
	}
	// default implementation
	return nil
}

func (m *MockExpenseRepository) GenerateNewId() (uint64, error) {
	m.GenerateNewIdCallCount++
	if m.GenerateNewIdFunc != nil {
		return m.GenerateNewIdFunc()
	}
	// default implementation
	return 1, nil
}

func (m *MockExpenseRepository) FindById(id uint64) (*domain.Expense, error) {
	m.FindByIdCallCount++
	if m.FindByIdFunc != nil {
		return m.FindByIdFunc(id)
	}
	// default implementation
	return &domain.Expense{
		Id:          1,
		Description: "Lunch",
		Amount:      20,
	}, nil
}
