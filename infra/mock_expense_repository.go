package infra

import (
	"github.com/t-sakoda/expense-tracker/domain"
)

type MockExpenseRepository struct {
	SaveFunc          func(expense *domain.Expense) error
	GenerateNewIdFunc func() (uint64, error)
	FindByIdFunc      func(id uint64) (*domain.Expense, error)
	DeleteFunc        func(id uint64) error
	FindAllFunc       func() ([]domain.Expense, error)

	SaveCallCount          int
	GenerateNewIdCallCount int
	FindByIdCallCount      int
	DeleteCallCount        int
	FindAllCallCount       int
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

func (m *MockExpenseRepository) Delete(id uint64) error {
	m.DeleteCallCount++
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	// default implementation
	return nil
}

func (m *MockExpenseRepository) FindAll() ([]domain.Expense, error) {
	m.FindAllCallCount++
	if m.FindAllFunc != nil {
		return m.FindAllFunc()
	}
	// default implementation
	return []domain.Expense{}, nil
}
