package domain

import (
	"fmt"
	"testing"
)

var userC = User{
	Name: "C",
	Id:   "C",
}

var event1 = Event{
	Users: []User{
		userA, userB, userC,
	},
	Id: "event1",
	Payments: []Payment{
		{
			Price: 1500,
			Payer: &userA,
			Payees: []User{
				userA, userB, userC,
			},
			Id: "1",
		},
		{
			Price: 1500,
			Payer: &userA,
			Payees: []User{
				userA, userB, userC,
			},
			Id: "1",
		},
	},
}

func TestExtractExchanges(t *testing.T) {
	tests := []struct {
		name      string
		input     Event
		exchanges []Exchange
	}{
		{
			name:  "event1を正しく計算できる",
			input: event1,
			exchanges: []Exchange{
				{
					Price: 1000,
					Payee: &userA,
					Payer: &userB,
				},
				{
					Price: 1000,
					Payee: &userA,
					Payer: &userC,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.PaymentSummaries().Exchanges()
			fmt.Println("\n\nexchanges:")
			for _, v := range got {
				fmt.Printf("%+v, payer: %v, payee: %v\n", v, v.Payer.Name, v.Payee.Name)
			}
			fmt.Println("")
			ok := exchangesAlike(got, tt.exchanges)
			if !ok {
				t.Errorf("expected %v, got %v", tt.exchanges, got)
			}
		})
	}
}

func exchangesAlike(exc1, exc2 []Exchange) bool {
	if len(exc1) != len(exc2) {
		return false
	}
	for _, v := range exc1 {
		ok := false
		for _, vv := range exc2 {
			if v.Price == vv.Price && v.Payee.Alike(*vv.Payee) && v.Payer.Alike(*vv.Payer) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
