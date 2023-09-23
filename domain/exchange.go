package domain

type Exchange struct {
	Price int

	// 精算時に支払う人
	Payee *User

	// 精算時に受け取る人
	Payer *User
}
