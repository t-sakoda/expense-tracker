package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/service"
)

func TestUpdateCmdRunE(t *testing.T) {
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
			cmd := &cobra.Command{}
			cmd.Flags().Uint64("id", tt.id, "")
			cmd.Flags().String("description", tt.description, "")
			cmd.Flags().Float64("amount", tt.amount, "")

			out := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(out)
			args := []string{}
			service := service.NewMockExpenseService()

			err := updateCmdRunE(cmd, args, service)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError {
				expected := "Expense updated successfully (ID: 1)\n"
				if out.String() != expected {
					t.Errorf("expected: %s, got: %s", expected, out.String())
				}
			}
		})
	}
}

func TestUpdateCmdRunEWithError(t *testing.T) {
	t.Run("when service.Update returns an error", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().Uint64("id", 1, "")
		cmd.Flags().String("description", "Lunch", "")
		cmd.Flags().Float64("amount", 20, "")

		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.UpdateFunc = func(id uint64, description string, amount float64) error {
			return errors.New("something went wrong")
		}

		err := updateCmdRunE(cmd, args, service)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		expected := "failed to update expense: something went wrong"
		if err.Error() != expected {
			t.Errorf("expected: %s, got: %s", expected, err.Error())
		}
	})
}
