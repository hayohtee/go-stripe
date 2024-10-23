package card

import (
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	StatusID       int
	Amount         int64
	Currency       string
	LastFourDigit  string
	BankReturnCode string
}

func (c *Card) Charge(currency string, amount int64) (*stripe.PaymentIntent, string, error) {
	return c.createPaymentIntent(currency, amount)
}

func (c *Card) createPaymentIntent(currency string, amount int64) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg string
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "your card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "your card is expired"
	default:
		msg = "your card was declined"
	}
	return msg
}
