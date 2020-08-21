package main

func moveZeroes(nums []int)  {
    cur := 0
    last := 0
    for ; cur < len(nums); cur ++ {
        if nums[cur] != 0 {
            tmp := nums[last]
            nums[last] = nums[cur]
            nums[cur] = tmp
            last ++
        }
    }
}