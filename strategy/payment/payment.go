package payment

type Payment interface {
	Pay()    //支付
	Refund() // 退款
}
