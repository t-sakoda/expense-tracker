package infra

import (
	"encoding/json"
	"testing"

	"github.com/spf13/afero"
	"github.com/t-sakoda/expense-tracker/domain"
)

func TestExpenseJsonRepositoryFindById(t *testing.T) {
	t.Run("when FindById is successful", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		file, err := fs.Create("test.json")
		if err != nil {
			t.Fatalf("failed to create file: %v", err)
		}
		defer file.Close()

		expense := &domain.Expense{
			Id:          1,
			Description: "test",
			Amount:      1000,
		}
		encoder := json.NewEncoder(file)
		if err := encoder.Encode([]domain.Expense{*expense}); err != nil {
			t.Fatalf("failed to encode JSON: %v", err)
		}

		repo := NewExpenseJsonRepository(file)
		actual, findErr := repo.FindById(1)
		if findErr != nil {
			t.Fatalf("failed to find expense: %v", findErr)
		}
		if *actual != *expense {
			t.Errorf("unexpected data: %v", actual)
		}
	})

	t.Run("when FindById fails", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		file, err := fs.Create("test.json")
		if err != nil {
			t.Fatalf("failed to create file: %v", err)
		}
		defer file.Close()

		repo := NewExpenseJsonRepository(file)
		actual, findErr := repo.FindById(1)
		if findErr == nil {
			t.Fatalf("unexpectedly succeeded: %v", actual)
		}
		if findErr.Error() != "expense not found: id=1" {
			t.Errorf("unexpected error: %v", findErr)
		}
	})
}
