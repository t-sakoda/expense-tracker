package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/service"
)

func TestAddCmdRunE(t *testing.T) {
	tests := []struct {
		description string
		amount      float64
		expectError bool
	}{
		{"Lunch", 20, false},
		{"Dinner", 50, false},
		{"", 20, true},
		{"Coffee", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("description", tt.description, "")
			cmd.Flags().Float64("amount", tt.amount, "")
			out := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(out)
			args := []string{}
			service := service.NewMockExpenseService()
			err := addCmdRunE(cmd, args, service)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError {
				expected := "Expense added successfully (ID: 1)\n"
				if out.String() != expected {
					t.Errorf("expected: %s, got: %s", expected, out.String())
				}
			}
		})
	}
}

func TestAddCmdRunEWithUseCaseError(t *testing.T) {
	tests := []struct {
		description string
		amount      float64
		expectError bool
	}{
		{"Lunch", 20, false},
		{"Dinner", 50, false},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("description", tt.description, "")
			cmd.Flags().Float64("amount", tt.amount, "")
			out := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(out)
			args := []string{}
			service := &service.MockExpenseService{}
			service.AddFunc = func(description string, amount float64) (uint64, error) {
				return 0, errors.New("something went wrong")
			}
			err := addCmdRunE(cmd, args, service)

			if err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
