package service

import (
	"errors"
	"testing"
	"time"

	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/infra"
)

func TestExpenseServiceAdd(t *testing.T) {
	tests := []struct {
		description string
		amount      float64
		expectError bool
	}{
		{"Lunch", 20, false},
		{"Dinner", 50, false},
		{"", 20, true},
		{"Lunch", -20, true},
	}

	clock := &infra.MockClock{}
	clock.NowFunc = func() time.Time {
		return time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := &infra.MockExpenseRepository{}
			s := NewExpenseService(repo, clock)
			id, err := s.Add(tt.description, tt.amount)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if tt.expectError {
				if id != 0 {
					t.Errorf("expected: 0, got: %d", id)
				}
				if repo.SaveCallCount != 0 {
					t.Errorf("expected: 0, got: %d", repo.SaveCallCount)
				}
			} else {
				if id == 0 {
					t.Errorf("expected: non-zero, got: %d", id)
				}
				if repo.SaveCallCount != 1 {
					t.Errorf("expected: 1, got: %d", repo.SaveCallCount)
				}
			}
		})
	}
}

func TestExpenseServiceAddWithError(t *testing.T) {
	clock := &infra.MockClock{}
	clock.NowFunc = func() time.Time {
		return time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	t.Run("Failed to generate id", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.GenerateNewIdFunc = func() (uint64, error) {
			return 0, errors.New("failed to generate id")
		}
		s := NewExpenseService(repo, clock)
		id, err := s.Add("Lunch", 20)

		if err == nil {
			t.Errorf("expected error: %v, got: %v", true, false)
		}
		if id != 0 {
			t.Errorf("expected: 0, got: %d", id)
		}
		if repo.GenerateNewIdCallCount != 1 {
			t.Errorf("expected: 1, got: %d", repo.GenerateNewIdCallCount)
		}
		if repo.SaveCallCount != 0 {
			t.Errorf("expected: 0, got: %d", repo.SaveCallCount)
		}
	})

	t.Run("Failed to save expense", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.SaveFunc = func(expense *domain.Expense) error {
			return errors.New("failed to save expense")
		}
		s := NewExpenseService(repo, clock)
		id, err := s.Add("Lunch", 20)

		if err == nil {
			t.Errorf("expected error: %v, got: %v", true, false)
		}
		if id != 0 {
			t.Errorf("expected: 0, got: %d", id)
		}
		if repo.GenerateNewIdCallCount != 1 {
			t.Errorf("expected: 1, got: %d", repo.GenerateNewIdCallCount)
		}
		if repo.SaveCallCount != 1 {
			t.Errorf("expected: 1, got: %d", repo.SaveCallCount)
		}
	})
}
