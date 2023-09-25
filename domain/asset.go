package domain

// Paymentから得た、支払った人とその支払の額
type Asset struct {
	// これは割り算が絡まないのでできるだけintで保持する
	Price int
	Owner *User
}

type AssetCollection []Asset

func (ac AssetCollection) AssetSum() int {
	sum := 0
	for _, v := range ac {
		sum += v.Price
	}
	return sum
}
