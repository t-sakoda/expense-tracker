package infra

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/spf13/afero"
)

func TestExpenseJsonRepositoryDelete(t *testing.T) {
	t.Run("when Delete is successful", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		file, err := fs.Create("test.json")
		if err != nil {
			t.Fatalf("failed to create file: %v", err)
		}
		defer file.Close()

		repo := NewExpenseJsonRepository(file)
		initialData := []map[string]interface{}{
			{
				"Id":          1,
				"Description": "test",
				"Amount":      1000,
				"Date":        "2021-01-01T12:34:56.789Z",
			},
		}
		encoder := json.NewEncoder(file)
		encodeErr := encoder.Encode(initialData)
		if encodeErr != nil {
			t.Fatalf("failed to encode initial data: %v", encodeErr)
		}

		deleteErr := repo.Delete(1)
		if deleteErr != nil {
			t.Fatalf("failed to delete expense: %v", deleteErr)
		}

		var actual []map[string]interface{}
		buffer := new(bytes.Buffer)
		file.Seek(0, io.SeekStart)
		buffer.ReadFrom(file)
		unmarshalErr := json.Unmarshal(buffer.Bytes(), &actual)
		if unmarshalErr != nil {
			t.Fatalf("failed to unmarshal JSON: %v", unmarshalErr)
		}
		if len(actual) != 0 {
			t.Errorf("unexpected saved data: %v", actual)
		}
	})

	t.Run("when Delete fails", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		file, err := fs.Create("test.json")
		if err != nil {
			t.Fatalf("failed to create file: %v", err)
		}

		repo := NewExpenseJsonRepository(file)
		initialData := []map[string]interface{}{
			{
				"Id":          1,
				"Description": "test",
				"Amount":      1000,
				"Date":        "2021-01-01T12:34:56.789Z",
			},
		}
		encoder := json.NewEncoder(file)
		encodeErr := encoder.Encode(initialData)
		if encodeErr != nil {
			t.Fatalf("failed to encode initial data: %v", encodeErr)
		}
		// Close the file to simulate a write error
		file.Close()

		deleteErr := repo.Delete(1)
		if deleteErr == nil {
			t.Fatalf("expected an error but got nil")
		}
	})
}
