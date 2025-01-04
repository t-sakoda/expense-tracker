package cmd

import (
	"bytes"
	"errors"
	"fmt"
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

func TestSummaryCmdRunEWithMonth(t *testing.T) {
	monthNameTests := []struct {
		monthName string
		month     int
	}{
		{"January", 1},
		{"February", 2},
		{"March", 3},
		{"April", 4},
		{"May", 5},
		{"June", 6},
		{"July", 7},
		{"August", 8},
		{"September", 9},
		{"October", 10},
		{"November", 11},
		{"December", 12},
	}

	for _, tt := range monthNameTests {
		t.Run("service.summaryMonth returns the total expenses", func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().Int("month", tt.month, "")
			out := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(out)
			args := []string{}
			service := &service.MockExpenseService{}
			service.SummaryMonthFunc = func(month int) (float64, error) {
				return 100.0, nil
			}

			err := summaryCmdRunE(cmd, args, service)
			if err != nil {
				t.Errorf("expected nil, got: %v", err)
			}

			expected := fmt.Sprintf("Total expenses for %s: $100.00\n", tt.monthName)
			if out.String() != expected {
				t.Errorf("expected: %s, got: %s", expected, out.String())
			}
		})

	}

	invalidMonthTests := []struct {
		month int
	}{
		{0},
		{13},
	}

	for _, tt := range invalidMonthTests {
		t.Run("invalid month flag is set", func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().Int("month", tt.month, "")
			out := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(out)
			args := []string{}
			service := &service.MockExpenseService{}
			service.SummaryMonthFunc = func(month int) (float64, error) {
				return 0, nil
			}

			err := summaryCmdRunE(cmd, args, service)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
