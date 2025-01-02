package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/t-sakoda/expense-tracker/domain"
)

type ExpenseJsonRepository struct {
	readWriteSeeker io.ReadWriteSeeker
	mutex           sync.Mutex
}

func NewExpenseJsonRepository(rws io.ReadWriteSeeker) *ExpenseJsonRepository {
	return &ExpenseJsonRepository{
		readWriteSeeker: rws,
	}
}

func (r *ExpenseJsonRepository) GenerateNewId() (uint64, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	expenses, err := r.readJson()
	if err != nil {
		return 0, fmt.Errorf("failed to read expenses: %w", err)
	}
	if len(expenses) == 0 {
		return 1, nil
	}

	var maxId uint64
	for _, e := range expenses {
		if e.Id > maxId {
			maxId = e.Id
		}
	}
	return maxId + 1, nil
}

func (r *ExpenseJsonRepository) FindAll() ([]domain.Expense, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.readJson()
}

func (r *ExpenseJsonRepository) Save(expense *domain.Expense) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	expenses, err := r.readJson()
	if err != nil {
		return fmt.Errorf("failed to find all expenses: %w", err)
	}

	var found bool = false
	for i, e := range expenses {
		if e.Id == expense.Id {
			expenses[i] = *expense
			found = true
			break
		}
	}
	if !found {
		expenses = append(expenses, *expense)
	}

	return r.writeJson(expenses)
}

func (r *ExpenseJsonRepository) readJson() ([]domain.Expense, error) {
	// Seek to the start of the file
	if _, err := r.readWriteSeeker.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek to the start of the file: %w", err)
	}

	var expenses []domain.Expense
	decoder := json.NewDecoder(r.readWriteSeeker)
	err := decoder.Decode(&expenses)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return expenses, nil
}

func (r *ExpenseJsonRepository) writeJson(expenses []domain.Expense) error {
	// Seek to the start of the file
	if _, err := r.readWriteSeeker.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to the start of the file: %w", err)
	}

	if err := r.clear(); err != nil {
		return fmt.Errorf("failed to clear the file: %w", err)
	}
	encoder := json.NewEncoder(r.readWriteSeeker)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(expenses); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	return nil
}

func (r *ExpenseJsonRepository) clear() error {
	if t, ok := r.readWriteSeeker.(interface{ Truncate(size int64) error }); ok {
		return t.Truncate(0) // Clear the file
	}
	return errors.New("readWriteSeeker does not support truncation")
}
