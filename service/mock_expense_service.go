package service

import (
	"errors"
)

/**
 * MockExpenseService
 */
type MockExpenseService struct {
}

func NewMockExpenseService() ExpenseServiceInterface {
	return &MockExpenseService{}
}

func (m *MockExpenseService) Add(description string, amount float64) (uint64, error) {
	return 1, nil
}

/**
 * MockExpenseServiceWithError
 */
type MockExpenseServiceWithError struct {
}

func NewMockExpenseServiceWithError() ExpenseServiceInterface {
	return &MockExpenseServiceWithError{}
}

func (m *MockExpenseServiceWithError) Add(description string, amount float64) (uint64, error) {
	return 0, errors.New("failed to save expense")
}
