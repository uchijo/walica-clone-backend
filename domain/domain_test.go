package domain

import (
	"math"
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

var event2 = Event{
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
		{
			Price: 1500,
			Payer: &userA,
			Payees: []User{
				userA, userB,
			},
			Id: "1",
		},
	},
}

// 詳細はどうあれ、整合性が保たれていれば良いテスト
func TestExtractExchangesRought(t *testing.T) {
	tests := []struct {
		name  string
		input Event
	}{
		{
			name:  "event1の整合性チェック",
			input: event1,
		},
		{
			name:  "event2の整合性チェック",
			input: event2,
		},
		{
			name: "旅行者二人、支払いなし",
			input: Event{
				Users:    []User{userA, userB},
				Id:       "a",
				Payments: []Payment{},
			},
		},
		{
			name: "旅行者二人、自分への支払いが1個",
			input: Event{
				Users: []User{userA, userB},
				Id:    "a",
				Payments: []Payment{
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA,
						},
						Id: "1",
					},
				},
			},
		},
		{
			name: "旅行者二人、自他への支払いが1個",
			input: Event{
				Users: []User{userA, userB},
				Id:    "a",
				Payments: []Payment{
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB,
						},
						Id: "1",
					},
				},
			},
		},
		{
			name: "Alice, Bob(e2eで死んだやつ)",
			input: Event{
				Users: []User{userA, userB},
				Id:    "a",
				Payments: []Payment{
					{
						Name: "buy ticket",
						Price: 3000,
						Payer: &userA,
						Payees: []User{
							userA, userB,
						},
						Id: "1",
					},
				},
			},
		},
		{
			name: "旅行者3人、割り切れない場合",
			input: Event{
				Users: []User{userA, userB, userC},
				Id:    "a",
				Payments: []Payment{
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB, userC,
						},
						Id: "1",
					},
				},
			},
		},
		{
			name: "旅行者3人、割り切れないのがたくさん",
			input: Event{
				Users: []User{userA, userB, userC},
				Id:    "a",
				Payments: []Payment{
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB, userC,
						},
						Id: "1",
					},
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB, userC,
						},
						Id: "1",
					},
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB, userC,
						},
						Id: "1",
					},
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB, userC,
						},
						Id: "1",
					},
					{
						Price: 1000,
						Payer: &userA,
						Payees: []User{
							userA, userB, userC,
						},
						Id: "1",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exc := tt.input.PaymentSummaries().Exchanges()
			for _, v := range tt.input.PaymentSummaries() {
				expectedTotal := v.Total()
				actualTotal := 0.0
				for _, vv := range exc {
					if vv.Payee.Alike(*v.User) {
						actualTotal += vv.price
					}
					if vv.Payer.Alike(*v.User) {
						actualTotal -= vv.price
					}
				}
				if math.Abs(actualTotal-expectedTotal) > 2 {
					t.Errorf("expected total is %v, but got %v", expectedTotal, actualTotal)
				}
			}
		})
	}
}

// より詳細なテスト。exchangeの中身まで気にする。
func TestExtractExchangesDetailed(t *testing.T) {
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
					price: 1000,
					Payee: &userA,
					Payer: &userB,
				},
				{
					price: 1000,
					Payee: &userA,
					Payer: &userC,
				},
			},
		},
		{
			name:  "event2を正しく計算できる",
			input: event2,
			exchanges: []Exchange{
				{
					price: 1750,
					Payee: &userA,
					Payer: &userB,
				},
				{
					price: 1000,
					Payee: &userA,
					Payer: &userC,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.PaymentSummaries().Exchanges()
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
			if v.price == vv.price && v.Payee.Alike(*vv.Payee) && v.Payer.Alike(*vv.Payer) {
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
