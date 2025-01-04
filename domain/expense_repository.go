package domain

type ExpenseRepository interface {
	GenerateNewId() (uint64, error)
	Save(expense *Expense) error
	FindById(id uint64) (*Expense, error)
	Delete(id uint64) error
}
