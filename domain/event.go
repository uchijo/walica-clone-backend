package domain

type Event struct {
	Users    UserCollection
	Name     string
	Id       string
	Payments PaymentCollection
}

func (e Event) PaymentSummaries() PaymentSummaryCollection {
	summaries := []PaymentSummary{}
	for i, v := range e.Users {
		summary := PaymentSummary{
			User:   &e.Users[i],
			Debts:  e.Payments.ExtractDebts(v),
			Assets: e.Payments.ExtractAssets(v),
		}
		summaries = append(summaries, summary)
	}
	return summaries
}
