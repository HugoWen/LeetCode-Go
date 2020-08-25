package main

func singleNumber(nums []int) int {
    one := 0
    two := 0
    for _, num := range nums {
        one = one ^ num & ^two
        two = two ^ num & ^one
    }
    return one
}