package domain

type Event struct {
	Users    UserCollection
	Id       string
	Payments PaymentCollection
}

func (e Event) PaymentSummaries() []PaymentSummary {
	summaries := []PaymentSummary{}
	for _, v := range e.Users {
		summary := PaymentSummary{
			User:   &v,
			Debts:  e.Payments.ExtractDebts(v),
			Assets: e.Payments.ExtractAssets(v),
		}
		summaries = append(summaries, summary)
	}
	return summaries
}
