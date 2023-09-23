package domain

import "testing"

// Exchange.Priceの小数点周りの処理をテスト
func TestExchangeRound(t *testing.T) {
	tests := []struct {
		name     string
		input    Exchange
		expected int
	}{
		{
			name: "入力が0",
			input: Exchange{
				Payer: &userA,
				Payee: &userB,
				price: 0.0,
			},
			expected: 0,
		},
		{
			name: "入力が1",
			input: Exchange{
				Payer: &userA,
				Payee: &userB,
				price: 1.0,
			},
			expected: 1,
		},
		{
			name: "入力が1.1のとき、1として出力される",
			input: Exchange{
				Payer: &userA,
				Payee: &userB,
				price: 1.1,
			},
			expected: 1,
		},
		{
			name: "入力が1.9のとき、2として出力される",
			input: Exchange{
				Payer: &userA,
				Payee: &userB,
				price: 1.9,
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Price()
			if got != tt.expected {
				t.Errorf("expected %v, but got %v. raw value was %v.", tt.expected, got, tt.input.price)
			}
		})
	}
}
