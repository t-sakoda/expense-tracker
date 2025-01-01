package domain

type Expense struct {
	Id          uint64
	Amount      float64
	Description string
}

func NewExpense(description string, amount float64) *Expense {
	return &Expense{
		Id:          0,
		Amount:      amount,
		Description: description,
	}
}
