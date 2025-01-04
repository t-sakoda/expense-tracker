package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

func updateCmdRunE(cmd *cobra.Command, _ []string, svc service.ExpenseServiceInterface) error {
	id, errId := cmd.Flags().GetUint64("id")
	amount, errA := cmd.Flags().GetFloat64("amount")
	description, errD := cmd.Flags().GetString("description")

	if errId != nil {
		return fmt.Errorf("failed to get ID: %w", errId)
	}
	if id == 0 {
		return fmt.Errorf("invalid expense ID: %d", id)
	}

	if errA != nil {
		return fmt.Errorf("failed to get amount: %w", errA)
	}
	if amount <= 0 {
		return fmt.Errorf("invalid expense amount: %f", amount)
	}

	if errD != nil {
		return fmt.Errorf("failed to get description: %w", errD)
	}
	if description == "" {
		return fmt.Errorf("description is required")
	}

	errUpdate := svc.Update(id, description, amount)
	if errUpdate != nil {
		return fmt.Errorf("failed to update expense: %w", errUpdate)
	}

	cmd.Printf("Expense updated successfully (ID: %d)\n", id)
	return nil
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an expense with a description and amount.",
	Example: `expense-tracker update --description "Dinner" --amount 50`,

	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		repo := infra.NewExpenseJsonRepository(file)
		service := service.NewExpenseService(repo)
		return updateCmdRunE(cmd, args, service)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Uint64("id", 0, "ID of the expense to update")
	updateCmd.Flags().String("description", "", "Description of the expense")
	updateCmd.Flags().Float64("amount", 0, "Amount of the expense")

	updateCmd.MarkFlagRequired("id")
}
