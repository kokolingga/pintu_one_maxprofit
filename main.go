package main

import (
	"fmt"
)

/*
Example 1:
Price of bitcoin each hour
Hour 1: $5
Hour 2: $4
Hour 3: $3
Hour 4: $2
Hour 5: $1
In this scenario, the maximum profit Jacky can make is: 0

Example 2:
Price of bitcoin each hour
Hour 1: $3
Hour 2: $2
Hour 3: $1
Hour 4: $5
Hour 5: $6
Hour 6: $2
In this scenario, the maximum profit Jacky can make is: $5, he will buy at 3rd hour and sell at 5th hour.
*/

type decision int

const (
	buying decision = iota
	selling
)

type profitTracker struct {
	lowestBuyingHour    int
	lowestBuyingPrice   int64
	highestSellingHour  int
	highestSellingPrice int64
	maxProfit           int64
}

func newProfitTracker() *profitTracker {
	return &profitTracker{}
}

func getMaxProfit(prices []int64, pt *profitTracker) int64 {
	if len(prices) == 0 {
		return 0
	}

	// (initial state)
	// as an initial step in decision making,
	// let's use the first element of the slice.
	pt.lowestBuyingPrice = prices[0]
	fmt.Println("(initial price checking)")
	fmt.Printf("%s : %v\n\n", "price", pt.lowestBuyingPrice)

	for index := 1; index < len(prices); index++ {
		if prices[index] < pt.lowestBuyingPrice {
			fmt.Printf("%20s : %12v\n", "current lowest price", pt.lowestBuyingPrice)
			fmt.Printf("%20s : %12v\n", "next hour price", prices[index])
			fmt.Printf("%s\n", "[it's cheaper. let's buy]")
			fmt.Printf("%s\n\n", "[marked as the lowest price]")
			pt.lowestBuyingPrice = prices[index]
			pt.lowestBuyingHour = index
		}

		if prices[index] > pt.lowestBuyingPrice {
			fmt.Printf("%20s : %12v\n", "current lowest price", pt.lowestBuyingPrice)
			fmt.Printf("%20s : %12v\n", "next hour price", prices[index])
			fmt.Printf("%s\n", "[a raise. should we sell?]")

			profit_amount := prices[index] - pt.lowestBuyingPrice
			fmt.Printf("%20s : %12v\n\n", "estimated profit", profit_amount)

			if profit_amount > pt.maxProfit {
				pt.maxProfit = profit_amount
				pt.highestSellingHour = index
				pt.highestSellingPrice = prices[index]
			}
		}
	}

	return pt.maxProfit
}

func main() {
	var prices = []int64{3, 2, 1, 5, 6, 2}
	// var prices = []int64{5, 4, 3, 2, 1}
	// var prices = []int64{137147048, 102162326, 199268418, 198975474, 253639272, 356694498, 225661554, 315177788, 328486079, 337443096, 279363057}

	pt := newProfitTracker()

	maxProfit := getMaxProfit(prices, pt)

	fmt.Printf("cheapest Buying : Hour[%v] : %v\n", pt.lowestBuyingHour, pt.lowestBuyingPrice)
	fmt.Printf("highest Selling : Hour[%v] : %v\n", pt.highestSellingHour, pt.highestSellingPrice)
	fmt.Printf("maxProfit : %v\n", pt.maxProfit)
	fmt.Printf("maxProfit : %v\n", maxProfit)
}
