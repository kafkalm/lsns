package main

import (
	"fmt"
)

func main() {
	cities, _ := LoadCityFromJSON("city.json")
	p, orders := CalBestPerFTGProfitRoute(500, cities, true)
	fmt.Println(p, orders)
	p, orders = CalBestTotalProfitRoute(500, cities, true)
	fmt.Println(p, orders)
	p, orders = CalBestReplenishProfitRoute(30, cities, true, 5)
	fmt.Println(p, orders)
	
}