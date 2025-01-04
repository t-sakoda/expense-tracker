package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

func addCmdRunE(cmd *cobra.Command, _ []string, svc service.ExpenseServiceInterface) error {
	amount, errA := cmd.Flags().GetFloat64("amount")
	description, errD := cmd.Flags().GetString("description")

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

	id, err := svc.Add(description, amount)
	if err != nil {
		return fmt.Errorf("failed to add expense: %w", err)
	}

	cmd.Printf("Expense added successfully (ID: %d)\n", id)
	return nil
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add an expense with a description and amount.",
	Example: `expense-tracker add --description "Lunch" --amount 20`,

	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		repo := infra.NewExpenseJsonRepository(file)
		clock := &infra.Clock{}

		service := service.NewExpenseService(repo, clock)
		return addCmdRunE(cmd, args, service)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("description", "", "Description of the expense")
	addCmd.Flags().Float64("amount", 0, "Amount of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
