package domain

type ExpenseRepository interface {
	GenerateNewId() (uint64, error)
	Save(expense *Expense) error
}
