package domain

type Payment struct {
	Price  int
	Payer  *User
	Payees []*User
	Id     string
}

type PaymentCollection []Payment
