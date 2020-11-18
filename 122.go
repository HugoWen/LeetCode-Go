package main

func maxProfit(prices []int) int {
    i := 0
    n := len(prices)
    amount := 0
    for ;i < n - 1; i ++ {
        if prices[i + 1] < prices[i] {
            continue
        } 
        amount += prices[i + 1] - prices[i]
    }
    return amount
}