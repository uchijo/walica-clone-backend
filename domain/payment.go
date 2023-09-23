package domain

type Payment struct {
	Price  int
	Payer  *User
	Payees UserCollection
	Id     string
}

type PaymentCollection []Payment

func (p Payment) Debt(u User) Debt {
	divider := p.Payees.Len()
	pricePerPerson := p.Price / divider
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
