package data

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/uchijo/walica-clone-backend/data/model"

	// "github.com/uchijo/walica-clone-backend/usecase"
	"github.com/uchijo/walica-clone-backend/util"
)

type RepositoryImpl struct{}

// var _ usecase.Repository = (*RepositoryImpl)(nil)

func (r RepositoryImpl) CreateEvent(name string, users []string) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("failed to generate error")
	}

	// create event
	event := model.Event{Name: name, ID: id}
	if err = util.DB.Create(&event).Error; err != nil {
		return "", err
	}

	// create users and associate them with event
	for _, v := range users {
		user := model.User{Name: v, EventId: id.String()}
		util.DB.Create(&user)
	}

	return id.String(), nil
}

func (r RepositoryImpl) CreatePayment(eventId, name, payer string, payees []string, price int) (string, error) {
	users := []model.User{}
	if err := util.DB.Where("ID in (?)", payees).Find(&users).Error; err != nil {
		return "", err
	}

	payerId, err := strconv.ParseUint(payer, 10, 64)
	if err != nil {
		return "", err
	}

	// create payment record
	payment := model.Payment{
		EventId: eventId,
		PayerId: uint(payerId),
		Price:   price,
		Name:    name,
		Payees:  users,
	}

	if err := util.DB.Create(&payment).Error; err != nil {
		return "", err
	}

	return strconv.FormatUint(uint64(payment.ID), 10), nil
}
