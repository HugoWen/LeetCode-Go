package main

func singleNumber(nums []int) int {
    ret := nums[0]
    if len(nums) > 1 {
        for i := 1; i < len(nums); i++ {
            ret = ret ^ nums[i]
        }
    }
    return ret
}