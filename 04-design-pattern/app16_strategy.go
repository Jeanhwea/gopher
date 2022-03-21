////////////////////////////////////////////////////////////////////////////////
// 策略模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 定义现金和银行卡两种支付策略
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

////////////////////////////////////////////////////////////////////////////////
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)

}

func main() {
	var payment *Payment
	payment = NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()

	payment = NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}

////////////////////////////////////////////////////////////////////////////////
// Pay $123 to Ada by cash
// Pay $888 to Bob by bank account 0002
