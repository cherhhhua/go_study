package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	fmt.Println(ans)
}

func maxProfit(prices []int) int {
	minprice := maxPrice(prices)
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}

func maxPrice(prices []int) int {
	result := prices[0]
	for _, price := range prices {
		if price > result {
			result = price
		}
	}
	return result
}
