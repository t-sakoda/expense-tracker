package cmd

import (
	"bytes"
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
