package infra

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/t-sakoda/expense-tracker/domain"
)

var mockDateStr = "2021-01-01T12:34:56.789Z"
var mockDate, _ = time.Parse(time.RFC3339, mockDateStr)

func TestExpenseJsonRepositorySave(t *testing.T) {
	t.Run("when Save is successful", func(t *testing.T) {
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
			Date:        mockDate,
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

		expected := []map[string]interface{}{
			{
				"Id":          uint64(1),
				"Description": "test",
				"Amount":      1000.0,
				"Date":        mockDateStr,
			},
		}
		if unmarshalErr != nil {
			t.Fatalf("failed to unmarshal JSON: %v", unmarshalErr)
		}
		if reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected saved data: %v", actual)
		}
	})

	t.Run("when Save fails", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		file, err := fs.Create("test.json")
		if err != nil {
			t.Fatalf("failed to create file: %v", err)
		}
		// Close the file to simulate a write error
		file.Close()

		repo := NewExpenseJsonRepository(file)
		expense := &domain.Expense{
			Id:          1,
			Description: "test",
			Amount:      1000,
			Date:        mockDate,
		}
		saveErr := repo.Save(expense)
		if saveErr == nil {
			t.Fatalf("expected an error but got nil")
		}
	})

	t.Run("when update the existing expense", func(t *testing.T) {
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
			Date:        mockDate,
		}
		initialData := []domain.Expense{*expense}
		encoder := json.NewEncoder(file)
		encodeErr := encoder.Encode(initialData)
		if encodeErr != nil {
			t.Fatalf("failed to encode initial data: %v", encodeErr)
		}

		repo := NewExpenseJsonRepository(file)
		updatedExpense := &domain.Expense{
			Id:          1,
			Description: "updated",
			Amount:      2000,
			Date:        mockDate,
		}
		saveErr := repo.Save(updatedExpense)
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

		expected := []map[string]interface{}{
			{
				"Id":          uint64(1),
				"Description": "updated",
				"Amount":      2000.0,
				"Date":        mockDateStr,
			},
		}
		if reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected saved data: %v", actual)
		}
	})
}
