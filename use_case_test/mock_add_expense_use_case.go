package use_case_test

import "fmt"

type MockAddExpenseUseCase struct{}

func (uc *MockAddExpenseUseCase) Execute(description string, amount float64) (uint64, error) {
	return 1, nil
}

type MockAddExpenseUseCaseWithError struct{}

func (uc *MockAddExpenseUseCaseWithError) Execute(description string, amount float64) (uint64, error) {
	return 0, fmt.Errorf("intentional error")
}
