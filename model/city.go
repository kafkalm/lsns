package model

var (
	FTG = map[Route]int{
		{FromCity: "City_1", ToCity: "City_2"}: 10,
		{FromCity: "City_2", ToCity: "City_1"}: 10,
	}
)

type Route struct {
	FromCity string
	ToCity string
}

type City struct {
	Name string `json:"name"`
	LocalProductNames []string `json:"local_product_names"`
	Inventory Inventory `json:"inventory"`
}

func (c City) FTG(anotherCity City) int {
	return FTG[Route{FromCity: c.Name, ToCity: anotherCity.Name}]
}

func (c City) GetLocalInventory() Inventory {
	return c.Inventory.Filter(func(productName string) bool {
		for _, localProductName := range c.LocalProductNames {
			if localProductName == productName {
				return true
			}
		}
		return false})
}

func (c City) CalMaxProfit(trainMaxLoad int, sellCity City) (profit int, orders []Order) {
	return c.GetLocalInventory().CalMaxProfit(trainMaxLoad, sellCity.Inventory)
}

// Replenish 补货
func (c City) Replenish(times int) {
	if times == 0 {
		return
	}
	for _, productName := range c.LocalProductNames {
		if times > 0 {
			c.Inventory.MultiProduct(c.Inventory.ProductInfo[productName], times)
		} else {
			c.Inventory.DivideProduct(c.Inventory.ProductInfo[productName], -times)
		}
	}
}