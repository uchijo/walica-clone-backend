package domain

type Debt struct {
	Price  int
	Debtor *User
}

type DebtCollection []Debt

func (dc DebtCollection) DebtSum() int {
	sum := 0
	for _, v := range dc {
		sum += v.Price
	}
	return sum
}
