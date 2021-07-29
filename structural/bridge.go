package structural

import (
	"fmt"
	"strings"
)

type Direction int

const (
	Buy Direction = iota
	Sell
)

// Hierarchy of Order direction types
type OrderType interface {
	Execute(o Order)
}

type BuyOrder struct {
	OrderDirection Direction
}

type SellOrder struct {
	OrderDirection Direction
}

//bridge
func (d *BuyOrder) Execute(o Order) {
	fmt.Println(d.OrderDirection, "Order Executed", o.String())
}

func (d *SellOrder) Execute(o Order) {
	fmt.Println(d.OrderDirection, "Order Executed", o.String())
}

// Hierarchy of order types
type Order struct {
	Type     OrderType // aggregation of order direction
	Symbol   string
	Quantity int
}

func NewOrder(d OrderType, s string, q int) Order {
	return Order{d, s, q}
}

func (o *Order) PlaceOrder() {
	fmt.Println("order placed: ", o.String())
	o.Type.Execute(*o)
}

func (o *Order) String() string {
	sb := strings.Builder{}
	sb.WriteString("Order Type: Market order\n")
	sb.WriteString(o.string())
	return sb.String()
}

func (o *Order) string() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Symbol: %s\n", o.Symbol))
	sb.WriteString(fmt.Sprintf("Quantity: %d\n", o.Quantity))
	return sb.String()
}

type LimitOrder struct {
	Order
	Price float32
}

func NewLimitOrder(d OrderType, s string, q int, p float32) LimitOrder {
	return LimitOrder{NewOrder(d, s, q), p}
}

func (o *LimitOrder) String() string {
	sb := strings.Builder{}
	sb.WriteString("Order Type: Limit order\n")
	sb.WriteString(o.string())
	sb.WriteString(fmt.Sprintf("Price: %f\n", o.Price))
	return sb.String()
}

type AllOrNoneOrder struct {
	Order
	IsFilled bool
}

func NewAllOrNoneOrder(d OrderType, s string, q int, isFilled bool) AllOrNoneOrder {
	return AllOrNoneOrder{NewOrder(d, s, q), isFilled}
}

func (o *AllOrNoneOrder) String() string {
	sb := strings.Builder{}
	sb.WriteString("Order Type: All-or-None order\n")
	sb.WriteString(o.string())
	sb.WriteString(fmt.Sprintf("IsFilled: %t\n", o.IsFilled))
	return sb.String()
}
