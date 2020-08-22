package main

func fib(n int) int {
    if n < 2 {
        return n
    }
    slice := make([]int, n + 1)
    slice[1] = 1
    for i := 2; i <= n; i ++ {
        slice[i] = slice[i - 1] + slice[i - 2]
    }
    return slice[n]
}