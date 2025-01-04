package service

import "github.com/t-sakoda/expense-tracker/domain"

/**
 * MockExpenseService
 */
type MockExpenseService struct {
	AddFunc     func(description string, amount float64) (uint64, error)
	UpdateFunc  func(id uint64, description string, amount float64) error
	DeleteFunc  func(id uint64) error
	ListFunc    func() ([]domain.Expense, error)
	SummaryFunc func() (float64, error)
}

func NewMockExpenseService() ExpenseServiceInterface {
	return &MockExpenseService{}
}

func (m *MockExpenseService) Add(description string, amount float64) (uint64, error) {
	if m.AddFunc != nil {
		return m.AddFunc(description, amount)
	}
	return 1, nil
}

func (m *MockExpenseService) Update(id uint64, description string, amount float64) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, description, amount)
	}
	return nil
}

func (m *MockExpenseService) Delete(id uint64) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func (m *MockExpenseService) List() ([]domain.Expense, error) {
	if m.ListFunc != nil {
		return m.ListFunc()
	}
	return []domain.Expense{}, nil
}

func (m *MockExpenseService) Summary() (float64, error) {
	if m.SummaryFunc != nil {
		return m.SummaryFunc()
	}
	return 0, nil
}
