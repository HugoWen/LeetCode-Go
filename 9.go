package main

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    origin := x
    revert := 0
    for x != 0 {
        revert = revert * 10 + x % 10
        x /= 10
    }

    return revert == origin
}