package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i ++ {
        dp[i] = make([]int, n)
    }

    for i := m - 1; i >= 0; i -- {
        for j := n - 1; j >= 0; j -- {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
                continue
            }

            if i == m - 1 && j == n - 1 {
                dp[i][j] = 1
            } else if i == m - 1 {
                dp[i][j] = dp[i][j + 1]
            } else if j == n - 1 {
                dp[i][j] = dp[i + 1][j]
            } else {
                dp[i][j] = dp[i + 1][j] + dp[i][j + 1]
            }
        }
    }
    return dp[0][0]
}