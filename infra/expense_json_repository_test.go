package infra

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/t-sakoda/expense-tracker/domain"
)

func TestExpenseJsonRepositorySave(t *testing.T) {
	fs := afero.NewMemMapFs()
	file, err := fs.Create("test.json")
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	repo := NewExpenseJsonRepository(file)
	expense := &domain.Expense{
		Id:          1,
		Description: "test",
		Amount:      1000,
	}
	saveErr := repo.Save(expense)
	if saveErr != nil {
		t.Fatalf("failed to save expense: %v", saveErr)
	}

	var actual []domain.Expense
	buffer := new(bytes.Buffer)
	file.Seek(0, io.SeekStart)
	buffer.ReadFrom(file)
	unmarshalErr := json.Unmarshal(buffer.Bytes(), &actual)
	if unmarshalErr != nil {
		t.Fatalf("failed to unmarshal JSON: %v", unmarshalErr)
	}
	if len(actual) != 1 || actual[0] != *expense {
		t.Errorf("unexpected saved data: %v", actual)
	}
}

func TestExpenseJsonRepositorySaveError(t *testing.T) {
	fs := new(afero.MemMapFs)
	file, err := afero.TempFile(fs, "", "ioutil-test")
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	file.Close()

	repo := NewExpenseJsonRepository(file)
	expense := &domain.Expense{
		Id:          1,
		Description: "test",
		Amount:      1000,
	}
	saveErr := repo.Save(expense)
	if saveErr == nil {
		t.Fatalf("expected an error but got nil")
	}
}
