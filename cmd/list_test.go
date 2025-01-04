package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/service"
)

func TestListCmdRunE(t *testing.T) {
	t.Run("service.List returns error", func(t *testing.T) {
		cmd := &cobra.Command{}
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.ListFunc = func() ([]*domain.Expense, error) {
			return nil, errors.New("something went wrong")
		}
		err := listCmdRunE(cmd, args, service)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("service.List returns expenses", func(t *testing.T) {
		cmd := &cobra.Command{}
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.ListFunc = func() ([]*domain.Expense, error) {
			return []*domain.Expense{
				{Id: 1, Description: "Lunch", Amount: 20},
				{Id: 2, Description: "Dinner", Amount: 50},
			}, nil
		}
		err := listCmdRunE(cmd, args, service)
		if err != nil {
			t.Errorf("expected nil, got: %v", err)
		}

		expected := "ID\tDate\tDescription\tAmount"
		expected += "1\tLunch\t$20.00\n"
		expected += "2\tDinner\t$50.00\n"
		if out.String() != expected {
			t.Errorf("expected: %s, got: %s", expected, out.String())
		}
	})
}
