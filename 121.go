package main

func maxProfit(prices []int) int {
    n := len(prices)
    if n <= 1 {
        return 0
    }
    i := 0
    j := 1
    pro := 0

    for j < n {
        if (prices[i] <= prices[j]) {
            pro = max(pro, prices[j] - prices[i])
            j ++
        } else {
            i = j
            j ++
        }
    }
    return pro
}

func max(i int, j int) int {
    if i >= j {
        return i
    }
    return j
}