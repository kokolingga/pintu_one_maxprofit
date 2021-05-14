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

type profit struct {
	lowestHour  int
	lowestPrice int64
	myDecision  decision
}

type decision int

const (
	buying decision = iota
	selling
)

func get_max_profit(prices []int64) profit {
	if len(prices) == 0 {
		return profit{}
	}

	// lowest price timing
	// for initial value, let's take the first element of the slice.
	lowest_buying_price := prices[0]
	lowest_buying_hour := 0
	myDecision := buying

	// fmt.Println("Initial value : ")
	// fmt.Println("first checking Index : ", lowest_buying_hour)
	fmt.Println("first checking Price : ", lowest_buying_price)
	fmt.Println("....................")

	for index := 1; index < len(prices); index++ {
		// index start from 1 (skip the first element)
		// fmt.Println("current Index : ", index)
		fmt.Println("current Price : ", prices[index])

		if myDecision == buying {
			if prices[index] < lowest_buying_price && index > lowest_buying_hour {
				fmt.Println("ah, it's cheaper. let's buy.")
				lowest_buying_price = prices[index]
				lowest_buying_hour = index
			}

			if prices[index] > lowest_buying_price {
				fmt.Printf("[lowest price : %v] a raise. should we sell?\n", lowest_buying_price)
			}
		}
		fmt.Println("....................")

	}

	return profit{lowestPrice: lowest_buying_price, lowestHour: lowest_buying_hour, myDecision: myDecision}
}

func main() {
	// var prices = []int64{3, 2, 1, 5, 6, 2}
	// var prices = []int64{5, 4, 3, 2, 1}
	var prices = []int64{137147048, 102162326, 199268418, 198975474, 253639272, 356694498, 225661554, 315177788, 328486079, 337443096, 279363057}

	profitData := get_max_profit(prices)

	fmt.Println("Lowest Price : ", profitData.lowestPrice)
	fmt.Println("Lowest Index : ", profitData.lowestHour)
	fmt.Println("My Decision : ", profitData.myDecision)
}
