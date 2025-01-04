package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/service"
)

func TestDeleteCmdRunE(t *testing.T) {
	tests := []struct {
		testDesc    string
		id          uint64
		expectError bool
	}{
		{"Valid param", 1, false},
		{"Invalid param", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.testDesc, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().Uint64("id", tt.id, "")

			out := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(out)
			args := []string{}
			service := service.NewMockExpenseService()

			err := deleteCmdRunE(cmd, args, service)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError {
				expected := "Expense deleted successfully"
				if out.String() != expected {
					t.Errorf("expected: %s, got: %s", expected, out.String())
				}
			}
		})
	}
}
