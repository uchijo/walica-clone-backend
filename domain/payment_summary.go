package domain

import (
	"errors"
	"math"
)

type PaymentSummary struct {
	User   *User
	Assets AssetCollection
	Debts  DebtCollection
}

type PaymentSummaryCollection []PaymentSummary

func (p PaymentSummary) tmpSummary() tmpSummary {
	return tmpSummary{
		user:  p.User,
		total: p.Total(),
	}
}

func (p PaymentSummary) Total() int {
	return p.Assets.AssetSum() - p.Debts.DebtSum()
}

func (p PaymentSummary) TotalAbs() int {
	return int(math.Abs(float64(p.Total())))
}

// func (c PaymentSummaryCollection) Exchanges() []Exchange {
//
// }

type tmpSummary struct {
	user  *User
	total int
}

func (ts tmpSummary) done() bool {
	return ts.total == 0
}

func (ts *tmpSummary) resolve(subject *tmpSummary) (Exchange, error) {
	// どっちもマイナスの場合、どっちもプラスの場合、精算済みの場合を弾く
	if ts.total*subject.total >= 0 {
		return Exchange{}, errors.New("invalid resolve")
	}

	// 負債と債権が全く同じ場合
	if ts.total+subject.total == 0 {
		ts.total = 0
		subject.total = 0
		return Exchange{
			Price: ts.total,
			Payee: bigger(ts, subject).user,
			Payer: smaller(ts, subject).user,
		}, nil
	}

	// tsが打ち消される場合
	if math.Abs(float64(ts.total)) < math.Abs(float64(subject.total)) {
		subject.total = ts.total + subject.total
		price := int(math.Abs(float64(ts.total)))
		ts.total = 0
		return Exchange{
			Price: price,
			Payee: bigger(ts, subject).user,
			Payer: smaller(ts, subject).user,
		}, nil
	}

	// 上以外、つまりsubjectが打ち消される場合
	ts.total = ts.total + subject.total
	price := int(math.Abs(float64(subject.total)))
	subject.total = 0
	return Exchange{
		Price: price,
		Payee: bigger(ts, subject).user,
		Payer: smaller(ts, subject).user,
	}, nil
}

func bigger(a, b *tmpSummary) *tmpSummary {
	if a.total > b.total {
		return a
	} else {
		return b
	}
}

func smaller(a, b *tmpSummary) *tmpSummary {
	if b.total > a.total {
		return a
	} else {
		return b
	}
}
