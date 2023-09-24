package domain

import "testing"

func TestPaymentAssetExtract(t *testing.T) {
	tests := []struct {
		name     string
		payments PaymentCollection
		user     User
		expected int
	}{
		{
			name: "alice, bob, buying 3000 ticket",
			payments: []Payment{
				{
					Name:  "buy ticket",
					Id:    "1",
					Price: 3000,
					Payer: &User{Name: "Alice", Id: "1"},
					Payees: []User{
						{Name: "Alice", Id: "1"},
						{Name: "Bob", Id: "2"},
					},
				},
			},
			user:     User{Name: "Alice", Id: "1"},
			expected: 3000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assets := tt.payments.ExtractAssets(tt.user)
			if assets.AssetSum() != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, assets.AssetSum())
			}
		})
	}
}
