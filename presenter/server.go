package presenter

import (
	"context"

	"github.com/uchijo/walica-clone-backend/domain"
	apipb "github.com/uchijo/walica-clone-backend/proto/proto/api"
	"github.com/uchijo/walica-clone-backend/usecase"
)

type server struct {
	apipb.UnimplementedWalicaCloneApiServer
	repositry *usecase.Repository
}

func NewServer(r usecase.Repository) *server {
	return &server{repositry: &r}
}

func (s *server) CreateEvent(ctx context.Context, req *apipb.CreateEventRequest) (*apipb.CreateEventReply, error) {
	eventId, err := usecase.CreateEvent(*s.repositry, req.Members, req.Name)
	if err != nil {
		return nil, err
	}
	return &apipb.CreateEventReply{Id: eventId}, nil
}

func (s *server) AddPayment(ctx context.Context, req *apipb.AddPaymentRequest) (*apipb.AddPaymentReply, error) {
	paymentId, err := usecase.AddPayment(
		*s.repositry,
		req.Name,
		req.EventId,
		req.PayerId,
		req.PayeeIds,
		int(req.Price),
	)
	if err != nil {
		return nil, err
	}
	return &apipb.AddPaymentReply{Id: paymentId}, nil
}

func (s *server) ReadInfo(ctx context.Context, req *apipb.ReadInfoRequest) (*apipb.ReadInfoReply, error) {
	out, err := usecase.ReadInfo(*s.repositry, req.Id)
	if err != nil {
		return nil, err
	}

	payments := []*apipb.Payment{}
	for _, v := range out.Payments {
		payments = append(payments, convertPayment(v))
	}

	exchanges := []*apipb.Exchange{}
	for _, v := range out.Exchanges {
		exchanges = append(exchanges, convertExchange(v))
	}

	summaries := []*apipb.PaymentSummary{}
	for _, v := range out.Summaries {
		summaries = append(summaries, convertSummary(v))
	}

	return &apipb.ReadInfoReply{
		Payments:     payments,
		Exchanges:    exchanges,
		Summaries:    summaries,
		TotalExpense: int32(out.TotalExpense),
		EventName:    out.EventName,
	}, nil
}

func (s *server) UpdatePayment(ctx context.Context, req *apipb.UpdatePaymentRequest) (*apipb.UpdatePaymentReply, error) {
	id, err := usecase.UpdatePayment(
		*s.repositry,
		req.PaymentId,
		req.Name,
		req.PayerId,
		req.PayeeIds,
		int(req.Price),
	)
	if err != nil {
		return nil, err
	}
	return &apipb.UpdatePaymentReply{
		PaymentId: id,
	}, nil
}

func (s *server) ReadAllUsers(ctx context.Context, req *apipb.ReadAllUsersRequest) (*apipb.ReadAllUsersReply, error) {
	users, err := usecase.ReadUsers(*s.repositry, req.EventId)
	if err != nil {
		return nil, err
	}
	apiUsers := []*apipb.User{}
	for _, v := range users {
		apiUsers = append(apiUsers, convertUser(v))
	}
	return &apipb.ReadAllUsersReply{
		Users: apiUsers,
	}, nil
}

func (s *server) ReadPayment(ctx context.Context, req *apipb.ReadPaymentRequest) (*apipb.ReadPaymentReply, error) {
	payment, err := usecase.ReadPayment(*s.repositry, req.PaymentId)
	if err != nil {
		return nil, err
	}
	return &apipb.ReadPaymentReply{
		Payment: convertPayment(payment),
	}, nil
}

func (s *server) DeletePayment(ctx context.Context, req *apipb.DeletePaymentRequest) (*apipb.DeletePaymentReply, error) {
	err := usecase.DeletePayment(*s.repositry, req.PaymentId)
	if err != nil {
		return nil, err
	}
	return &apipb.DeletePaymentReply{
		PaymentId: req.PaymentId,
	}, nil
}

func convertSummary(s domain.PaymentSummary) *apipb.PaymentSummary {
	return &apipb.PaymentSummary{
		User:         convertUser(*s.User),
		TotalExpense: int32(s.Debts.DebtSum()),
	}
}

func convertExchange(e domain.Exchange) *apipb.Exchange {
	return &apipb.Exchange{
		Price: int32(e.Price()),
		Payer: convertUser(*e.Payer),
		Payee: convertUser(*e.Payee),
	}
}

func convertPayment(p domain.Payment) *apipb.Payment {
	payees := []*apipb.User{}
	for _, v := range p.Payees {
		payees = append(payees, convertUser(v))
	}
	return &apipb.Payment{
		Name:   p.Name,
		Id:     p.Id,
		Price:  int32(p.Price),
		Payer:  convertUser(*p.Payer),
		Payees: payees,
	}
}

func convertUser(u domain.User) *apipb.User {
	return &apipb.User{
		Name: u.Name,
		Id:   u.Id,
	}
}
