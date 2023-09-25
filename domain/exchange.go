package domain

import "math"

// 精算時、誰が誰にいくら払うかを表現するオブジェクト
type Exchange struct {
	// 整合性のためできるだけfloatで持っておく
	price float64

	// 精算時に支払う人
	Payer *User

	// 精算時に受け取る人
	Payee *User
}

// 四捨五入して返す
func (e Exchange) Price() int {
	return int(math.Round(e.price))
}
