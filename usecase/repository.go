package usecase

import "github.com/uchijo/walica-clone-backend/domain"

type Repository interface {
	// CreateEvent creates event.
	// input
	//   - name: event's name
	//   - users: users name
	// output
	//   - event's id
	//   - error
	CreateEvent(name string, users []string) (string, error)

	// ReadAllPayments returns all payments associated with given event id
	// input
	//   - eventId: event's id
	// output
	//   - list of payment info
	//   - error
	// ReadAllPayments(eventId string) (domain.PaymentCollection, error)

	// ReadPayment returns a payment with given payment id
	// input
	//   - paymentId: id of payment
	// output
	//   - payment info
	//   - error
	// ReadPayment(paymentId string) domain.Payment

	// CreatePayment creates payment record
	// input
	//   - eventId: id of the event
	//   - name: name of this payment
	//   - payer: payer's user id
	//   - payees: lit of payees' user id
	//   - price: price of this payment
	// output
	//   - payment's id
	//   - error
	CreatePayment(eventId, name, payer string, payees []string, price int) (string, error)

	// UpdatePayment updates payment record
	// input
	//   - id of updating payment
	//   - name: name of this payment
	//   - payer: payer's user id
	//   - payees: lit of payees' user id
	//   - price: price of this payment
	// output
	//   - payment's id
	//   - error
	UpdatePayment(paymentId, name, payer string, payees []string, price int) (string, error)

	// DeletePayment deletes payment record
	// input
	//   - id of deleting payment
	// output
	//   - error
	// DeletePayment(paymentId string) error

	// ReadAllUsers reads all users in the event
	// input
	//   - id of the event
	// output
	//   - list of users
	//   - error
	ReadAllUsers(eventId string) (domain.UserCollection, error)

	// ReadEventInfo reads event information by id
	// input
	//   - id of the event
	// output
	//   - event
	//   - error
	ReadEventInfo(eventId string) (domain.Event, error)
}
