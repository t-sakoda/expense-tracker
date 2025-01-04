package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary of all expenses",
	Example: `expense-tracker summary
expense-tracker summary --month 8`,

	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		repo := infra.NewExpenseJsonRepository(file)
		clock := &infra.Clock{}

		svc := service.NewExpenseService(repo, clock)
		return summaryCmdRunE(cmd, args, svc)
	},
}

func summaryCmdRunE(cmd *cobra.Command, _ []string, svc service.ExpenseServiceInterface) error {
	month, errMonth := cmd.Flags().GetUint8("month")
	// month flag is not set
	if errMonth != nil || month == 0 {
		total, err := svc.Summary()
		if err != nil {
			return fmt.Errorf("failed to get summary: %w", err)
		}
		cmd.Println(fmt.Sprintf("Total expenses: $%.2f", total))
		return nil
	}
	// valid month flag is set
	if 1 <= month && month <= 12 {
		total, err := svc.SummaryMonth(month)
		if err != nil {
			return fmt.Errorf("failed to get summary: %w", err)
		}
		monthName := time.Month(month).String()
		cmd.Println(fmt.Sprintf("Total expenses for %s: $%.2f", monthName, total))
		return nil
	}
	// invalid month flag is set
	return errors.New("invalid month")
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().Uint8("month", 0, "Month to summarize of current year")
}
