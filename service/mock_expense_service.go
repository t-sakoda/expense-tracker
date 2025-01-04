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

func (m *MockExpenseService) Update(id uint64, description string, amount float64) error {
	return nil
}

func (m *MockExpenseService) Delete(id uint64) error {
	return nil
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

func (m *MockExpenseServiceWithError) Update(id uint64, description string, amount float64) error {
	return errors.New("failed to update expense")
}

func (m *MockExpenseServiceWithError) Delete(id uint64) error {
	return errors.New("failed to delete expense")
}
