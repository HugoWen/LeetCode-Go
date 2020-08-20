package main

func majorityElement(nums []int) int {
    countMap := make(map[int]int)
    numLen := len(nums)
    for i := 0; i < numLen; i ++ {
        countMap[nums[i]] ++
        if countMap[nums[i]] > numLen / 2 {
            return nums[i]
        }
    }
    return nums[0]
}