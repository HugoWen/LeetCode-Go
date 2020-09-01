package main

func predictTheWinner(nums []int) bool {
    l := len(nums)
    dp := make([][]int, l)
    for i := 0; i < l; i ++ {
        dp[i] = make([]int, l)
        dp[i][i] = nums[i]
    }

    var max func(int, int) int
    max = func(num1 int, num2 int) int {
        if num1 >= num2 {
            return num1
        }
        return num2
    }

    for i := l - 2; i >= 0; i -- {
        for j := i + 1; j <= l - 1; j ++ {
            dp[i][j] = max(nums[i] - dp[i + 1][j], nums[j] - dp[i][j - 1])
        } 
    }

    return dp[0][l - 1] >= 0
}