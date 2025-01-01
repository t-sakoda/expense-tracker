package use_case

import (
	"testing"

	"github.com/t-sakoda/expense-tracker/infra_test"
)

func TestAddExpenseUseCaseExecute(t *testing.T) {
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
			repo := &infra_test.MockExpenseRepository{}
			uc := &AddExpenseUseCase{
				Repo: repo,
			}
			id, err := uc.Execute(tt.description, tt.amount)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError {
				if id != 1 {
					t.Errorf("expected: 1, got: %d", id)
				}
			}

			// Check if Insert method is called
			if tt.expectError {
				if repo.InsertCallCount() != 0 {
					t.Errorf("expected: 0, got: %d", repo.InsertCallCount())
				}
			} else {
				if repo.InsertCallCount() != 1 {
					t.Errorf("expected: 1, got: %d", repo.InsertCallCount())
				}
			}
		})
	}
}
