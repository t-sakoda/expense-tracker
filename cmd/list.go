package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

func listCmdRunE(cmd *cobra.Command, _ []string, svc service.ExpenseServiceInterface) error {
	expenses, err := svc.List()
	if err != nil {
		return fmt.Errorf("failed to list expenses: %w", err)
	}

	// Print the expenses
	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 1, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintf(w, "ID\tDate\tDescription\tAmount\n")
	for _, expense := range expenses {
		date := expense.Date.Format("2006-01-02")
		fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n", expense.Id, date, expense.Description, expense.Amount)
	}

	return nil
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all expenses",
	Example: `expense-tracker list`,

	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		repo := infra.NewExpenseJsonRepository(file)
		clock := &infra.Clock{}
		svc := service.NewExpenseService(repo, clock)
		return listCmdRunE(cmd, args, svc)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
