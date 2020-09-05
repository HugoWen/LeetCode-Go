package main

import "strconv"

func getPermutation(n int, k int) string {
    x, y := jiecheng(n, k)
    s := strconv.Itoa(x)
    for i := 1; i <= n; i ++ {
        if i == x {
            continue;
        }
    }
}

func jiecheng(n int, k int) int, int {
    if k == 1 {
        return 1
    }
    dp := []int{}
    dp[1] = 0
    dp[2] = 1
    for i := 3; i <= n; i ++ {
        dp[i] = dp[i - 1] * (i - 1)
    }

    x := k / dp[n]
    y := k % dp[n]
    if y != 0 {
        x ++
    }

    return x, k - dp[n] * (x - 1)
}