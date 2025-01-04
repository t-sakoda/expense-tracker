package service

import (
	"errors"
	"testing"
	"time"

	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/infra"
)

func TestExpenseServiceSummaryMonth(t *testing.T) {
	expenses := []domain.Expense{
		{Id: 1, Description: "expense 1", Amount: 10.0, Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Id: 2, Description: "expense 2", Amount: 20.0, Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
		{Id: 3, Description: "expense 3", Amount: 30.0, Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
		{Id: 4, Description: "expense 4", Amount: 40.0, Date: time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)},
		{Id: 5, Description: "expense 5", Amount: 50.0, Date: time.Date(2022, 3, 2, 0, 0, 0, 0, time.UTC)},
	}

	tests := []struct {
		now            time.Time
		expectedAmount float64
	}{
		{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 60.0},
		{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC), 40.0},
		{time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC), 50.0},
		{time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC), 0.0},
	}

	for _, tt := range tests {
		t.Run("FindAll returns expenses", func(t *testing.T) {
			repo := &infra.MockExpenseRepository{}
			repo.FindAllFunc = func() ([]domain.Expense, error) {
				return expenses, nil
			}
			clock := &infra.MockClock{}
			clock.NowFunc = func() time.Time {
				return tt.now
			}
			svc := NewExpenseService(repo, clock)

			total, err := svc.SummaryMonth(uint8(tt.now.Month()))
			if err != nil {
				t.Errorf("expected nil, got: %v", err)
			}

			expected := tt.expectedAmount
			if total != expected {
				t.Errorf("expected: %.2f, got: %.2f", expected, total)
			}
		})
	}

	t.Run("FindAll returns error", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.FindAllFunc = func() ([]domain.Expense, error) {
			return nil, errors.New("something went wrong")
		}
		clock := &infra.MockClock{}
		svc := NewExpenseService(repo, clock)

		_, err := svc.SummaryMonth(1)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
