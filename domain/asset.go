package domain

type Asset struct {
	Price int
	Owner *User
}

type AssetCollection []Asset
