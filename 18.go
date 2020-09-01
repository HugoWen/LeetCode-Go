package main

import "sort"

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return [][]int{}
    }

    ret := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums) - 3; i ++ {
        if i != 0 && nums[i] == nums[i - 1] {
            continue
        }
        for j := i + 1; j < len(nums) - 2; j ++ {
            if j != i + 1 && nums[j] == nums[j - 1] {
                continue
            }
            left := j + 1
            right := len(nums) - 1
            for left < right {
                // fmt.Println(i, j, left, right, ret)
                if nums[i] + nums[j] + nums[left] + nums[right] == target {
                    ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
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
                } else if nums[i] + nums[j] + nums[left] + nums[right] < target {
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