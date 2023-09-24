package data

import (
	"errors"

	"github.com/google/uuid"
	"github.com/uchijo/walica-clone-backend/data/model"
	"github.com/uchijo/walica-clone-backend/util"
)

type RepositoryImpl struct{}

// var _ usecase.Repository = (*RepositoryImpl)(nil)

func (r RepositoryImpl) CreateEvent(name string, users []string) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("failed to generate error")
	}

	if err = util.DB.Create(&model.Event{Name: name, ID: id}).Error; err != nil {
		return "", err
	}

	return id.String(), nil
}
