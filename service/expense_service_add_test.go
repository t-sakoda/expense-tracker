package service

import (
	"testing"

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

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := &infra.MockExpenseRepository{}
			s := NewExpenseService(repo)
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
	repo := &infra.MockExpenseRepositoryWithError{}
	s := NewExpenseService(repo)
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
}
