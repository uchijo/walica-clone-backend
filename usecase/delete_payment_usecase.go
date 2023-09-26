package usecase

func DeletePayment(repo Repository, paymentId string) error {
	return repo.DeletePayment(paymentId)
}
