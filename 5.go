package main

func longestPalindrome(s string) string {
    l := len(s)
    if l < 2 {
        return s
    }

    dp := make([][]bool, l)
    for i := 0; i < l; i++ {
        dp[i] = make([]bool, l)
    }
    max := 1
    begin := 0

    for j := 1; j < l; j ++ {
        for i := 0; i < j; i ++ {
            if i == j {
                dp[i][j] = true
                continue
            }
            if s[i] == s[j] {
                if j - i < 3 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i + 1][j - 1]
                }
            } else {
                dp[i][j] = false
            }

            if dp[i][j] == true && j - i + 1 > max {
                max = j - i + 1
                begin = i
            }
        }
    } 
    return s[begin : begin + max]
}