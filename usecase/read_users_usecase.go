package usecase

import "github.com/uchijo/walica-clone-backend/domain"

func ReadUsers(repository Repository, eventId string) ([]domain.User, error) {
	return repository.ReadAllUsers(eventId)
}
