package domain

// 支払い履歴
type Payment struct {
	Name   string
	Price  int
	Payer  *User
	Payees UserCollection
	Id     string
}

// 支払い履歴のリスト
type PaymentCollection []Payment

func (p Payment) Debt(u User) Debt {
	divider := float64(p.Payees.Len())
	pricePerPerson := float64(p.Price) / divider
	return Debt{
		Price:  pricePerPerson,
		Debtor: &u,
	}
}

func (pc PaymentCollection) ExtractDebts(u User) DebtCollection {
	debts := []Debt{}
	for _, v := range pc {
		if v.Payees.Contains(u) {
			debts = append(debts, v.Debt(u))
		}
	}
	return debts
}

func (p Payment) Asset(u User) Asset {
	return Asset{
		Price: p.Price,
		Owner: &u,
	}
}

func (pc PaymentCollection) ExtractAssets(u User) AssetCollection {
	assets := []Asset{}
	for _, v := range pc {
		if v.Payer.Alike(u) {
			assets = append(assets, v.Asset(u))
		}
	}
	return assets
}

func (pc PaymentCollection) PaymentSum() int {
	sum := 0
	for _, v := range pc {
		sum += v.Price
	}
	return sum
}
