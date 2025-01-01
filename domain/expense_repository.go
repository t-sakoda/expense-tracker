package domain

type ExpenseRepository interface {
	GenerateNewId() uint64
	Insert(expense *Expense) error
}
