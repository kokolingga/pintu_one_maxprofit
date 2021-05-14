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
}

func get_max_profit(prices []int64) profit {

	if len(prices) == 0 {
		return profit{}
	}

	// lowest price timing
	// for initial value, let's take the first element of the map.
	lowest_price := prices[0]
	lowest_hour := 0
	fmt.Println("Initial value : ")
	fmt.Println("first taking Price : ", lowest_price)
	fmt.Println("first taking Index : ", lowest_hour)
	fmt.Println("....................")

	for index := 1; index < len(prices); index++ {
		if prices[index] < lowest_price && index > lowest_hour {
			lowest_price = prices[index]
			lowest_hour = index
		}
	}

	return profit{lowestPrice: lowest_price, lowestHour: lowest_hour}
}

func main() {
	var prices = []int64{3, 2, 1, 5, 6, 2}

	profitData := get_max_profit(prices)

	fmt.Println("Lowest Price : ", profitData.lowestPrice)
	fmt.Println("Lowest Index : ", profitData.lowestHour)

}
