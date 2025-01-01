package domain

type ExpenseRepository interface {
	Create(expense *Expense) error
}
