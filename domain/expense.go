package domain

type Expense struct {
	id          uint64
	amount      float64
	description string
}

func (e *Expense) Id() uint64 {
	return e.id
}
func (e *Expense) Amount() float64 {
	return e.amount
}
func (e *Expense) Description() string {
	return e.description
}

func NewExpense(id uint64, amount float64, description string) *Expense {
	return &Expense{
		id:          id,
		amount:      amount,
		description: description,
	}
}
