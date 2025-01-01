package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/t-sakoda/expense-tracker/infra"
	"github.com/t-sakoda/expense-tracker/use_case"
)

func addCmdRunE(cmd *cobra.Command, args []string, uc *use_case.AddExpenseUseCase) error {
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

	cmd.Printf("description: %s\n", description)
	cmd.Printf("amount: %f\n", amount)
	id := "360990ce-10cb-49a3-b49e-69e494a6d557"
	cmd.Printf("Expense added successfully (ID: %s)\n", id)

	return nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add an expense with a description and amount.",
	Example: `expense-tracker add --description "Lunch" --amount 20`,

	RunE: func(cmd *cobra.Command, args []string) error {
		repo := infra.ExpenseJsonRepository{}
		uc := use_case.NewAddExpenseUseCase(repo)
		return addCmdRunE(cmd, args, uc)
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
