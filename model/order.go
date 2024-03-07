package model

type Order struct {
	ProductName string `json:"product_name"`
	BuyPrice    int    `json:"buy_price"`
	SellPrice   int    `json:"sell_price"`
	Quantity    int    `json:"quantity"`
}

func (o Order) Profit() int {
	return (o.SellPrice - o.BuyPrice) * o.Quantity
}

func CalculateOrdersProfit(orders []Order) int {
	profit := 0
	for _, order := range orders {
		profit += order.Profit()
	}
	return profit
}
