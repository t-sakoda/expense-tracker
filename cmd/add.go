package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var description string
var amount float64

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense with a description and amount.",
	Long: `Add an expense with a description and amount. For example:

expense-tracker add --description "Lunch" --amount 20`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("description: ", description)
		fmt.Println("amount: ", amount)
		id := "360990ce-10cb-49a3-b49e-69e494a6d557"
		fmt.Printf("Expense added successfully (ID: %s)\n", id)
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
	addCmd.Flags().StringVar(&description, "description", "", "Description of the expense")
	addCmd.Flags().Float64Var(&amount, "amount", 0, "Amount of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
