package main

import "strings"
import "math"

func myAtoi(str string) int {
    res := 0
    sign := 1
    str = strings.TrimSpace(str)
    for i, v := range str {
        if v >= '0' && v <= '9' {
            res = res * 10 + int(v - '0')
        } else if v == '+' && i == 0 {
            sign = 1 
        } else if v == '-' && i == 0 {
            sign = -1
        } else {
            break
        }

        if res > math.MaxInt32 {
            if sign == -1 {
                return math.MinInt32
            } else {
                return math.MaxInt32
            }
        }
    }
    return sign * res
}