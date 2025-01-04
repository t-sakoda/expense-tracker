package service

import (
	"errors"
	"testing"

	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/infra"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		id          uint64
		description string
		amount      float64
		expectError bool
	}{
		{1, "Lunch", 20, false},
		{1, "Dinner", 50, false},
		{1, "", 20, true},
		{1, "Coffee", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := &infra.MockExpenseRepository{}
			s := NewExpenseService(repo)

			err := s.Update(tt.id, tt.description, tt.amount)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
		})
	}
}

func TestUpdateWithError(t *testing.T) {
	t.Run("Expense not found", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.FindByIdFunc = func(id uint64) (*domain.Expense, error) {
			return nil, errors.New("expense not found")
		}
		s := NewExpenseService(repo)
		err := s.Update(1, "Lunch", 20)
		if err != ErrExpenseNotFound {
			t.Errorf("expected: %v, got: %v", ErrExpenseNotFound, err)
		}
	})

	t.Run("Invalid description", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		s := NewExpenseService(repo)
		err := s.Update(1, "", 20)
		if err != ErrInvalidParameter {
			t.Errorf("expected: %v, got: %v", ErrInvalidParameter, err)
		}
	})

	t.Run("Invalid amount", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		s := NewExpenseService(repo)
		err := s.Update(1, "Lunch", 0)
		if err != ErrInvalidParameter {
			t.Errorf("expected: %v, got: %v", ErrInvalidParameter, err)
		}
	})
}
