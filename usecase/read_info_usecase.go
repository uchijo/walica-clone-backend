package usecase

import "github.com/uchijo/walica-clone-backend/domain"

func ReadInfo(repo Repository, eventId string) (ReadInfoOutput, error) {
	event, err := repo.ReadEventInfo(eventId)
	if err != nil {
		return ReadInfoOutput{}, nil
	}
	return ReadInfoOutput{
		Payments:     event.Payments,
		Exchanges:    event.PaymentSummaries().Exchanges(),
		Summaries:    event.PaymentSummaries(),
		TotalExpense: event.Payments.PaymentSum(),
		EventName:    event.Name,
	}, nil
}

type ReadInfoOutput struct {
	Payments     domain.PaymentCollection
	Exchanges    []domain.Exchange
	Summaries    domain.PaymentSummaryCollection
	TotalExpense int
	EventName    string
}
