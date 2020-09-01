package main

import "sort"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    ret := nums[0] + nums[1] + nums[2]

    var diff func(int, int) int
    diff = func(num1 int, num2 int) int {
        r := num1 - num2
        if r >= 0 {
            return r
        }
        return -r
    }

    for i := 0; i < len(nums) - 2; i ++ {
        // if nums[i] >= target {
        //     if diff(nums[i] + nums[i + 1] + nums[i + 2], target) < diff(ret, target) {
        //         return nums[i] + nums[i + 1] + nums[i + 2]
        //     }
        //     return ret
        // }
        
        if i == 0 || nums[i] != nums[i - 1] {
            left := i + 1
            right := len(nums) - 1
            for left < right {
                // fmt.Println(i, left, right, ret)
                cur := nums[i] + nums[left] + nums[right]
                if cur == target {
                    return target
                }
                if diff(cur, target) < diff(ret, target) {
                    ret = cur
                }
                if cur < target {
                    left ++
                    for left < right {
                        if nums[left] == nums[left - 1] {
                            left ++
                        } else {
                            break
                        }
                    }
                } else {
                    right --
                    for left < right {
                        if nums[right] == nums[right + 1] {
                            right --
                        } else {
                            break
                        }
                    }
                }
            }
        }
    }
    return ret
}