package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/service"
)

func TestSummaryCmdRunE(t *testing.T) {
	t.Run("service.summary returns the total expenses", func(t *testing.T) {
		cmd := &cobra.Command{}
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.SummaryFunc = func() (float64, error) {
			return 100.0, nil
		}

		err := summaryCmdRunE(cmd, args, service)
		if err != nil {
			t.Errorf("expected nil, got: %v", err)
		}

		expected := "Total expenses: $100.00\n"
		if out.String() != expected {
			t.Errorf("expected: %s, got: %s", expected, out.String())
		}
	})

	t.Run("service.summary returns an error", func(t *testing.T) {
		cmd := &cobra.Command{}
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		cmd.SetErr(out)
		args := []string{}
		service := &service.MockExpenseService{}
		service.SummaryFunc = func() (float64, error) {
			return 0, errors.New("something went wrong")
		}

		err := summaryCmdRunE(cmd, args, service)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
