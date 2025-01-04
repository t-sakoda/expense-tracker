package cmd

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/service"
)

var mockDate1 = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
var mockDate2 = time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)

func TestListCmdRunE(t *testing.T) {
	t.Run("service.List returns error", func(t *testing.T) {
		cmd := &cobra.Command{}
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.ListFunc = func() ([]domain.Expense, error) {
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
		service.ListFunc = func() ([]domain.Expense, error) {
			return []domain.Expense{
				{Id: 1, Description: "Lunch at the restaurant", Amount: 20, Date: mockDate1},
				{Id: 2, Description: "Dinner", Amount: 50, Date: mockDate2},
			}, nil
		}
		err := listCmdRunE(cmd, args, service)
		if err != nil {
			t.Errorf("expected nil, got: %v", err)
		}

		expected := "ID Date       Description             Amount\n"
		expected += "1  2021-01-01 Lunch at the restaurant $20.00\n"
		expected += "2  2021-01-02 Dinner                  $50.00\n"
		if out.String() != expected {
			t.Errorf("expected: %s, got: %s", expected, out.String())
		}
	})
}
