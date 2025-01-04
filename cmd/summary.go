package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

var summaryCmd = &cobra.Command{
	Use:     "summary",
	Short:   "Summary of all expenses",
	Example: `expense-tracker summary`,

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
	total, err := svc.Summary()
	if err != nil {
		return fmt.Errorf("failed to get summary: %w", err)
	}

	cmd.Println(fmt.Sprintf("Total expenses: $%.2f", total))
	return nil
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
