package service

import (
	"errors"
	"testing"
	"time"

	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/infra"
)

func TestExpenseServiceSummary(t *testing.T) {
	t.Run("repository.FindAll returns expenses", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.FindAllFunc = func() ([]domain.Expense, error) {
			return []domain.Expense{
				{Id: 1, Description: "expense 1", Amount: 10.0, Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Id: 2, Description: "expense 2", Amount: 20.0, Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Id: 3, Description: "expense 3", Amount: 30.0, Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
			}, nil
		}
		clock := &infra.MockClock{}
		svc := NewExpenseService(repo, clock)

		total, err := svc.Summary()
		if err != nil {
			t.Errorf("expected nil, got: %v", err)
		}

		expected := 60.0
		if total != expected {
			t.Errorf("expected: %.2f, got: %.2f", expected, total)
		}
	})

	t.Run("repository.FindAll returns an error", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.FindAllFunc = func() ([]domain.Expense, error) {
			return nil, errors.New("something went wrong")
		}
		clock := &infra.MockClock{}
		svc := NewExpenseService(repo, clock)

		_, err := svc.Summary()
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
