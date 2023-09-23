package usecase

func UpdatePayment(repo Repository, paymentId, name, payer string, payees []string, price int) (string, error) {
	return repo.UpdatePayment(paymentId, name, payer, payees, price)
}
