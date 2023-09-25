package domain

// Paymentから得た、支払ってもらった人と負債の額
type Debt struct {
	// 割り算が絡むのでfloatにして後で均す
	Price  float64
	Debtor *User
}

type DebtCollection []Debt

func (dc DebtCollection) DebtSum() float64 {
	sum := 0.0
	for _, v := range dc {
		sum += v.Price
	}
	return sum
}
