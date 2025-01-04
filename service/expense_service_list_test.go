package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/t-sakoda/expense-tracker/domain"
	"github.com/t-sakoda/expense-tracker/infra"
)

func TestExpenseServiceList(t *testing.T) {
	clock := &infra.MockClock{}
	clock.NowFunc = func() time.Time {
		return time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	t.Run("When the repository's FindAll method successfully returns a list of expenses", func(t *testing.T) {
		expected := []domain.Expense{
			{Id: 1, Description: "item1", Amount: 100, Date: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Id: 2, Description: "item2", Amount: 200, Date: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)},
		}
		repo := &infra.MockExpenseRepository{}
		repo.FindAllFunc = func() ([]domain.Expense, error) {
			return expected, nil
		}
		service := NewExpenseService(repo, clock)
		actual, err := service.List()
		if err != nil {
			t.Errorf("service.List() returned error: %v", err)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("service.List() returned %v, want %v", expected, actual)
		}
	})

	t.Run("When the repository's FindAll method returns an error", func(t *testing.T) {
		repo := &infra.MockExpenseRepository{}
		repo.FindAllFunc = func() ([]domain.Expense, error) {
			return nil, errors.New("something went wrong")
		}
		service := NewExpenseService(repo, clock)
		_, err := service.List()
		if err == nil {
			t.Error("service.List() did not return an error")
		}
	})
}
