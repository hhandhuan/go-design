package payment

import "fmt"

type WeChatPay struct{}

func (p *WeChatPay) Pay() {
	fmt.Println("微信支付")
}

func (p *WeChatPay) Refund() {
	fmt.Println("微信退款")
}
