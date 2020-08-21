package main

func climbStairs(n int) int {
    first := 0
    second := 0
    third := 1
    for i := 1; i <= n; i ++ {
        first = second
        second = third
        third = first + second
    }
    return third
}