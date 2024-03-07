package main

import (
	"lsns/model"
)

// CalBestReplenishProfitRoute
// 限制最大进货书最优解
func CalBestReplenishProfitRoute(trainMaxLoad int, cities []model.City, round bool, maxReplenishTimes int) (profit float64, cityOrders map[string][]model.Order) {
	maxProfit := float64(0)
	maxOrders := map[string][]model.Order{}
	for i := 1; i <= maxReplenishTimes; i++ {
		for j := range cities {
			cities[j].Replenish(i)
		}
		profit, orders := CalBestTotalProfitRoute(trainMaxLoad, cities, round)
		if profit > maxProfit {
			maxProfit = profit
			maxOrders = orders
		}
		for j := range cities {
			cities[j].Replenish(-i)
		}
	}
	return maxProfit, maxOrders
}

// CalBestPerFTGProfitRoute
// 单疲劳值最优解
func CalBestPerFTGProfitRoute(trainMaxLoad int, cities []model.City, round bool) (profit float64, cityOrders map[string][]model.Order) {
	return CalBestProfitRoute(trainMaxLoad, cities, round, func(profit int, fromCity, toCity model.City) float64 {
		if round {
			return float64(profit) / (float64(fromCity.FTG(toCity)) + float64(toCity.FTG(fromCity)))
		}
		return float64(profit) / float64(fromCity.FTG(toCity))
	})
}

// CalBestTotalProfitRoute
// 最高收益
func CalBestTotalProfitRoute(trainMaxLoad int, cities []model.City, round bool) (profit float64, cityOrders map[string][]model.Order) {
	return CalBestProfitRoute(trainMaxLoad, cities, round, func(profit int, fromCity, toCity model.City) float64 {
		return float64(profit)
	})
}

func CalBestProfitRoute(trainMaxLoad int, cities []model.City, round bool, profitWight func(profit int, fromCity, toCity model.City) float64) (profit float64, cityOrders map[string][]model.Order) {
	if profitWight == nil {
		profitWight = func(profit int, fromCity, toCity model.City) float64 {return float64(profit)}
	}
	maxProfit := float64(0)
	maxCityOrders := make(map[string][]model.Order)
	for i := 0; i < len(cities); i++ {
		for j := i+1; j < len(cities); j++ {
			i2jProfit, i2jOrders := cities[i].CalMaxProfit(trainMaxLoad, cities[j])
			j2iProfit, j2iOrders := cities[j].CalMaxProfit(trainMaxLoad, cities[i])
			if round {
				profit := profitWight(i2jProfit + j2iProfit, cities[i], cities[j])
				if profit > maxProfit {
					maxProfit = profit
					maxCityOrders = map[string][]model.Order{
						cities[i].Name: i2jOrders,
						cities[j].Name: j2iOrders,
					}
				}
			} else {
				profit := profitWight(i2jProfit, cities[i], cities[j])
				if  profit > maxProfit {
					maxProfit = profit
					maxCityOrders = map[string][]model.Order{
						cities[i].Name: i2jOrders,
					}
				}
				profit = profitWight(j2iProfit, cities[j], cities[i])
				if  profit > maxProfit {
					maxProfit = profit
					maxCityOrders = map[string][]model.Order{
						cities[j].Name: j2iOrders,
					}
				}
			}
		}
	}
	return maxProfit, maxCityOrders
}
