package model

import (
	"fmt"
)

type Order struct {
	ProductName string `json:"product_name"`
	BuyPrice    int    `json:"buy_price"`
	SellPrice   int    `json:"sell_price"`
	Quantity    int    `json:"quantity"`
}

func (o Order) Profit() int {
	return (o.SellPrice - o.BuyPrice) * o.Quantity
}

func (o Order) KeyInfo() string {
	return fmt.Sprintf("%s 买入: %d", o.ProductName, o.Quantity)
}

func CalculateOrdersProfit(orders []Order) int {
	profit := 0
	for _, order := range orders {
		profit += order.Profit()
	}
	return profit
}
