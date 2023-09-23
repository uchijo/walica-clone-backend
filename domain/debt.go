package domain

type Debt struct {
	Price  int
	Debtor *User
}

type DebtCollection []Debt
