package domain

type Asset struct {
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
