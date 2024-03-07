package model

import (
	"lsns/util"

	"golang.org/x/exp/slices"
)

type Inventory struct {
	ProductQuantity map[string]int     `json:"product_quantity"`
	ProductInfo     map[string]Product `json:"product_info"`
}

func (i Inventory) AddProduct(p Product, quantity int) {

}

func (i Inventory) MultiProduct(p Product, times int) {
	i.ProductQuantity[p.Name] *= times
	i.ProductInfo[p.Name] = p
}

func (i Inventory) DivideProduct(p Product, times int) {
	i.ProductQuantity[p.Name] /= times
	i.ProductInfo[p.Name] = p
}

func (i Inventory) Filter(f func(productName string) bool) Inventory {
	filteredInventory := Inventory{
		ProductQuantity: make(map[string]int),
		ProductInfo:     make(map[string]Product),
	}

	for productName, qty := range i.ProductQuantity {
		if f(productName) {
			filteredInventory.ProductQuantity[productName] = qty
		}
	}

	for productName, product := range i.ProductInfo {
		if f(productName) {
			filteredInventory.ProductInfo[productName] = product
		}
	}

	return filteredInventory
}

// func (i Inventory) Sub(another Inventory) Inventory {
// 	subInventory := Inventory{
// 		ProductQuantity: make(map[string]int),
// 		ProductInfo: make(map[string]Product),
// 	}
// 	for productName, qty := range i.ProductQuantity {
// 		subInventory.ProductQuantity[productName] = qty
// 		subInventory.ProductInfo[productName] = Product{
// 			Name: productName,
// 			Price: another.ProductInfo[productName].Price - i.ProductInfo[productName].Price,
// 		}
// 	}
// 	return subInventory
// }

func (i Inventory) CalMaxProfit(trainMaxLoad int, sellInventory Inventory) (profit int, orders []Order) {
	products := make([]Product, 0, len(i.ProductQuantity))
	for productName, qty := range i.ProductQuantity {
		if qty <= 0 {
			continue
		}
		products = append(products, i.ProductInfo[productName])
	}

	slices.SortFunc[[]Product](products, func(a, b Product) int {
		aProfit := sellInventory.ProductInfo[a.Name].Price - a.Price
		bProfit := sellInventory.ProductInfo[b.Name].Price - b.Price
		if aProfit < bProfit {
			return 1
		} else if aProfit > bProfit {
			return -1
		}
		return 0
	})

	curLoad := 0
	orders = []Order{}
	for _, product := range products {
		if curLoad >= trainMaxLoad {
			break
		}
		q := util.Min(i.ProductQuantity[product.Name], trainMaxLoad-curLoad)
		if q == 0 {
			continue
		}
		orders = append(orders, Order{
			ProductName: product.Name,
			BuyPrice:    product.Price,
			SellPrice:   sellInventory.ProductInfo[product.Name].Price,
			Quantity:    q,
		})
		curLoad += q
	}

	return CalculateOrdersProfit(orders), orders
}
