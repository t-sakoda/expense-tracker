package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/service"
)

const jsonFilePath = "expenses.json"

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

// addCmd represents the add command
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
		service := service.NewExpenseService(repo)
		return addCmdRunE(cmd, args, service)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().String("description", "", "Description of the expense")
	addCmd.Flags().Float64("amount", 0, "Amount of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
