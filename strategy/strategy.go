package strategy

import (
	"design/strategy/payment"
)

type Order struct {
	BizNo   string
	Amount  float64
	Payment payment.Payment
}

func CreateOrder(payment payment.Payment) *Order {
	return &Order{
		BizNo:   "XNSS123123213213",
		Amount:  22.22,
		Payment: payment,
	}
}

func (o *Order) SetPayment(payment payment.Payment) {
	o.Payment = payment
}

func (o *Order) Pay() {
	o.Payment.Pay()
}

func (o *Order) Refund() {
	o.Payment.Refund()
}

// 	order := strategy.CreateOrder(&payment.AliPay{})
//	order.Pay()
//	order.Refund()
//
//	order.SetPayment(&payment.WeChatPay{})
//	order.Pay()
//	order.Refund()
