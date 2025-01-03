package infra

import (
	"errors"

	"github.com/t-sakoda/expense-tracker/domain"
)

/**
 * MockExpenseRepository
 */
type MockExpenseRepository struct {
	SaveCallCount int
}

func (m *MockExpenseRepository) Save(expense *domain.Expense) error {
	m.SaveCallCount++
	return nil
}

func (m *MockExpenseRepository) GenerateNewId() (uint64, error) {
	return 1, nil
}

/**
 * MockExpenseRepositoryWithError
 */
type MockExpenseRepositoryWithError struct {
	SaveCallCount          int
	GenerateNewIdCallCount int
}

func (m *MockExpenseRepositoryWithError) Save(expense *domain.Expense) error {
	m.SaveCallCount++
	return errors.New("failed to save expense")
}

func (m *MockExpenseRepositoryWithError) GenerateNewId() (uint64, error) {
	m.GenerateNewIdCallCount++
	return 0, errors.New("failed to generate id")
}
