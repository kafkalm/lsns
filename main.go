package main

import (
	"fmt"
	"encoding/json"
	"lsns/model"
)

func main() {
	cities, _ := LoadCityFromJSON("city.json")
	// p, orders := CalBestPerFTGProfitRoute(500, cities, true)
	// fmt.Println(p, InlineJSON(orders), "\n")
	// PrintOrders(p, orders)
	p, orders := CalBestTotalProfitRoute(500, cities, true)
	// fmt.Println(p, InlineJSON(orders), "\n")
	PrintOrders(p, orders)
	fmt.Println("=====================================================")
	p, orders = CalBestReplenishProfitRoute(500, cities, true, 2)
	// fmt.Println(int(p), InlineJSON(orders), "\n")
	PrintOrders(p, orders)
}

func InlineJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func PrintOrders(p float64,  orders map[string][]model.Order) {
	for city, order := range orders {
		fmt.Println(city)
		for _, o := range order {
			fmt.Println(o.KeyInfo())
		}
		fmt.Println()
	}
	fmt.Println("利润：", int(p), "\n")
}