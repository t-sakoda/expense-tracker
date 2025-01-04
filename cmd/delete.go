package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

func deleteCmdRunE(cmd *cobra.Command, _ []string, svc service.ExpenseServiceInterface) error {
	id, errId := cmd.Flags().GetUint64("id")
	if errId != nil {
		return fmt.Errorf("failed to get ID: %w", errId)
	}
	if id == 0 {
		return fmt.Errorf("invalid ID")
	}

	if err := svc.Delete(id); err != nil {
		return fmt.Errorf("failed to delete expense: %w", err)
	}

	cmd.Printf("Expense deleted successfully")
	return nil
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an expense.",
	Example: `expense-tracker delete --id 1`,

	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		repo := infra.NewExpenseJsonRepository(file)
		service := service.NewExpenseService(repo)
		return deleteCmdRunE(cmd, args, service)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Uint64("id", 0, "ID of the expense")
	deleteCmd.MarkFlagRequired("id")
}
