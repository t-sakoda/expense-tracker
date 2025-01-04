package service

import (
	"errors"
	"testing"

	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/infra"
)

func TestDelete(t *testing.T) {
	t.Run("Expense exists", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		s := NewExpenseService(repo)

		err := s.Delete(1)
		if err != nil {
			t.Errorf("expected: nil, got: %v", err)
		}
	})

	t.Run("Expense not found", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.FindByIdFunc = func(id uint64) (*domain.Expense, error) {
			return nil, errors.New("expense not found")
		}
		s := NewExpenseService(repo)
		err := s.Delete(1)
		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})
}

func TestDeleteWithError(t *testing.T) {
	t.Run("When repo.Delete returns an error", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.DeleteFunc = func(id uint64) error {
			return errors.New("something went wrong")
		}
		s := NewExpenseService(repo)
		err := s.Delete(1)
		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})
}
