package usecase

func AddPayment(repo Repository, name, eventId, payer string, payees []string, price int) (string, error) {
	return repo.CreatePayment(eventId, name, payer, payees, price)
}
