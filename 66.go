package main

func plusOne(digits []int) []int {
    l := len(digits)
    for i := l - 1; i >= 0; i -- {
        digits[i]++
        if digits[i] == 10 {
            digits[i] = 0
        } else {
            return digits
        }
    }

    digits[0] = 1
    digits = append(digits, 0)
    return digits
}