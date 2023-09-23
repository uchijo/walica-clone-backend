package domain

import "errors"

type Event struct {
	Users    UserCollection
	Id       string
	Payments PaymentCollection
}

func (e Event) DebtForUser(u User) (DebtCollection, error) {
	if !e.Users.Contains(u) {
		return nil, errors.New("user not found")
	}

	return e.Payments.ExtractDebts(u), nil
}

func (e Event) AssetsForUser(u User) (AssetCollection, error) {
	if !e.Users.Contains(u) {
		return nil, errors.New("user not found")
	}

	return e.Payments.ExtractAssets(u), nil
}
