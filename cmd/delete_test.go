package cmd

import (
	"bytes"
	"errors"
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

func TestDeleteCmdRunEWithError(t *testing.T) {
	t.Run("when service.Delete returns an error", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().Uint64("id", 1, "")

		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.DeleteFunc = func(id uint64) error {
			return errors.New("something went wrong")
		}

		err := deleteCmdRunE(cmd, args, service)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		expected := "failed to delete expense: something went wrong"
		if err.Error() != expected {
			t.Errorf("expected: %s, got: %s", expected, err.Error())
		}
	})
}
