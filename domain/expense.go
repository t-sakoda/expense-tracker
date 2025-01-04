package domain

import "time"

type Expense struct {
	Id          uint64
	Description string
	Amount      float64
	Date        time.Time
}
