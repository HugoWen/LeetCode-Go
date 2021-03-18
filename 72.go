package main

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func min(x, y, z int) int {
	if x > y {
		if y > z {
			return z
		} else {
			return y
		}
	} else {
		if x > z {
			return z
		} else {
			return x
		}
	}
}
