package domain

type Exchange struct {
	Price int

	// 精算時に支払う人
	Payer *User

	// 精算時に受け取る人
	Payee *User
}
