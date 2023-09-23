package domain

type Exchange struct {
	Price int
	Payee *User
	Payer *User
}
