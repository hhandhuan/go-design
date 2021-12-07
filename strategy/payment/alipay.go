package payment

import "fmt"

type AliPay struct{}

func (p *AliPay) Pay() {
	fmt.Println("支付宝支付")
}

func (p *AliPay) Refund() {
	fmt.Println("支付宝退款")
}
