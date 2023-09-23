package domain

import (
	"errors"
	"math"
)

// ユーザとそれに紐づいたすべての負債と債権をまとめる
type PaymentSummary struct {
	User   *User
	Assets AssetCollection
	Debts  DebtCollection
}

// 各ユーザのすべての貸し借りのリスト
type PaymentSummaryCollection []PaymentSummary

func (p PaymentSummary) tmpSummary() tmpSummary {
	return tmpSummary{
		user:  p.User,
		total: p.Total(),
	}
}

func (p PaymentSummary) Total() float64 {
	return float64(p.Assets.AssetSum()) - p.Debts.DebtSum()
}

// 割り切れないときとかに数円の誤差が生まれる
func (c PaymentSummaryCollection) Exchanges() []Exchange {
	summaries := []tmpSummary{}
	for _, v := range c {
		summaries = append(summaries, v.tmpSummary())
	}

	exchanges := []Exchange{}

	for {
		if isDone(summaries) {
			break
		}

		for outer := range summaries {
			if summaries[outer].done() {
				continue
			}

			for inner := outer + 1; inner < len(summaries); inner++ {
				if summaries[inner].done() {
					continue
				}
				exc, err := summaries[outer].resolve(&summaries[inner])
				if err != nil {
					continue
				}
				exchanges = append(exchanges, exc)
				break
			}
		}
	}

	return exchanges
}

type tmpSummary struct {
	user  *User
	total float64
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
		price := math.Abs(ts.total)
		data := Exchange{
			price: price,
			Payee: bigger(ts, subject).user,
			Payer: smaller(ts, subject).user,
		}
		ts.total = 0
		subject.total = 0
		return data, nil
	}

	// tsが打ち消される場合
	if math.Abs(ts.total) < math.Abs(subject.total) {
		subject.total = ts.total + subject.total
		price := math.Abs(ts.total)
		ts.total = 0
		return Exchange{
			price: price,
			Payee: bigger(ts, subject).user,
			Payer: smaller(ts, subject).user,
		}, nil
	}

	// 上以外、つまりsubjectが打ち消される場合
	ts.total = ts.total + subject.total
	price := math.Abs(subject.total)
	subject.total = 0
	return Exchange{
		price: price,
		Payee: bigger(ts, subject).user,
		Payer: smaller(ts, subject).user,
	}, nil
}

// 支払い完了が1以下だったら全体の計算も終わり
func isDone(c []tmpSummary) bool {
	doneCount := 0
	for _, v := range c {
		if v.done() {
			doneCount += 1
		}
	}
	return len(c)-doneCount <= 1
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
