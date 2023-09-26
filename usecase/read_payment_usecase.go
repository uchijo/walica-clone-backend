package usecase

import "github.com/uchijo/walica-clone-backend/domain"

func ReadPayment(repository Repository, paymentId string) (domain.Payment, error) {
	return repository.ReadPayment(paymentId)
}
