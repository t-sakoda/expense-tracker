package domain

type ExpenseRepository interface {
	Insert(expense *Expense) error
}
