package data

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/uchijo/walica-clone-backend/data/model"
	"gorm.io/gorm"

	"github.com/uchijo/walica-clone-backend/domain"

	"github.com/uchijo/walica-clone-backend/usecase"
	"github.com/uchijo/walica-clone-backend/util"
)

type RepositoryImpl struct {
	usecase.Repository
}

var _ usecase.Repository = (*RepositoryImpl)(nil)

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

func (r RepositoryImpl) DeletePayment(paymentId string) error {
	if err := util.DB.Delete(&model.Payment{}, paymentId).Error; err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) UpdatePayment(paymentId, name, payer string, payees []string, price int) (string, error) {
	users := []model.User{}
	if err := util.DB.Where("ID in (?)", payees).Find(&users).Error; err != nil {
		return "", err
	}

	row := &model.Payment{}
	if err := util.DB.First(row, paymentId).Error; err != nil {
		return "", err
	}

	payerId, err := strconv.ParseUint(payer, 10, 64)
	if err != nil {
		return "", err
	}

	// 一応トランザクション
	err = util.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&row).Updates(model.Payment{
			Name:    name,
			EventId: row.EventId,
			Price:   price,
			PayerId: uint(payerId),
		}).Error; err != nil {
			return err
		}

		asocs := tx.Model(&row).Association("Payees")
		err = asocs.Replace(users)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return paymentId, nil
}

func (r RepositoryImpl) ReadEventInfo(eventId string) (domain.Event, error) {
	rawEvent := model.Event{}
	if err := util.DB.Where("id = ?", eventId).Preload("Payments").First(&rawEvent).Error; err != nil {
		return domain.Event{}, err
	}

	rawUsers := []model.User{}
	if err := util.DB.Where("event_id == ?", eventId).Find(&rawUsers).Error; err != nil {
		return domain.Event{}, err
	}
	users := []domain.User{}
	for _, v := range rawUsers {
		users = append(users, convertUser(v))
	}

	payments := []domain.Payment{}
	for _, v := range rawEvent.Payments {
		p, err := convertPayment(v)
		if err != nil {
			return domain.Event{}, err
		}
		payments = append(payments, p)
	}

	event := domain.Event{
		Id:       eventId,
		Name:     rawEvent.Name,
		Users:    users,
		Payments: payments,
	}

	return event, nil
}

func ReadAllUsers(eventId string) (domain.UserCollection, error) {
	rawUsers := []model.User{}
	if err := util.DB.Where("event_id == ?", eventId).Find(&rawUsers).Error; err != nil {
		return nil, err
	}
	users := []domain.User{}
	for _, v := range rawUsers {
		users = append(users, convertUser(v))
	}
	return users, nil
}

func convertUser(input model.User) domain.User {
	return domain.User{
		Name: input.Name,
		Id:   strconv.FormatUint(uint64(input.ID), 10),
	}
}

func convertPayment(input model.Payment) (domain.Payment, error) {
	rawPayer := model.User{}
	if err := util.DB.Where("id = ?", input.PayerId).First(&rawPayer).Error; err != nil {
		return domain.Payment{}, nil
	}
	payer := convertUser(rawPayer)
	result := &model.Payment{}
	if err := util.DB.
		Where("id = ?", input.ID).
		Preload("Payees").
		First(result).Error; err != nil {
		return domain.Payment{}, nil
	}

	payees := []domain.User{}
	for _, v := range result.Payees {
		payees = append(payees, convertUser(v))
	}
	return domain.Payment{
		Name:   input.Name,
		Price:  input.Price,
		Payer:  &payer,
		Payees: payees,
		Id:     strconv.FormatUint(uint64(input.ID), 10),
	}, nil
}
