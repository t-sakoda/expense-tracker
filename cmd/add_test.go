package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_addCmd(t *testing.T) {
	tests := []struct {
		description string
		amount      float64
		expectError bool
	}{
		{"Lunch", 20, false},
		{"Dinner", 50, false},
		{"", 20, true},
		{"Coffee", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			out := new(bytes.Buffer)
			rootCmd.SetOut(out)
			rootCmd.SetErr(out)

			args := []string{"add", "--description", tt.description, "--amount", fmt.Sprintf("%f", tt.amount)}
			rootCmd.SetArgs(args)

			err := rootCmd.Execute()

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError {
				expectedOutput := fmt.Sprintf("description: %s\namount: %f\nExpense added successfully (ID: 360990ce-10cb-49a3-b49e-69e494a6d557)\n", tt.description, tt.amount)
				if out.String() != expectedOutput {
					t.Errorf("expected output: %s, got: %s", expectedOutput, out.String())
				}
			}
		})
	}
}
