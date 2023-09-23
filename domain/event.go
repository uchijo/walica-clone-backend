package domain

type Event struct {
	Users    []*User
	Id       string
	Payments PaymentCollection
}
