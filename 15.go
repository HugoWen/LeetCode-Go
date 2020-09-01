package main

import "sort"

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return [][]int{}
    }
    ret := [][]int{}
    sort.Ints(nums)
    for i := 0; i <= len(nums) - 3; i ++ {
        if nums[i] > 0 {
            break;
        }
        if i == 0 || nums[i] != nums[i - 1] {
            left := i + 1
            right := len(nums) - 1
            for left < right {
                // fmt.Println(i, left, right, ret)
                if nums[i] + nums[left] + nums[right] == 0 {
                    ret = append(ret, []int{nums[i], nums[left], nums[right]})
                    left ++
                    right --
                    for left < right {
                        if nums[left] != nums[left - 1] && nums[right] != nums[right + 1] {
                            break
                        }
                        if nums[left] == nums[left - 1] {
                            left ++
                        }
                        if nums[right] == nums[right + 1] {
                            right --
                        }
                    }
                } else if nums[i] + nums[left] + nums[right] < 0 {
                    left ++
                    for left < right {
                        if nums[left] == nums[left - 1] {
                            left ++
                        } else {
                            break;
                        }
                    }
                } else {
                    right --
                    for right > left {
                        if nums[right] == nums[right + 1] {
                            right --
                        } else {
                            break;
                        }
                    }
                }
            }
        }
    }
    return ret
}