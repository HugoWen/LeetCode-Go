package main

import "math"

const (
    EPSILON = 1E-6
)

func judgePoint24(nums []int) bool {
    transNums := []float64{}
    for i := 0; i < len(nums); i ++ {
        transNums = append(transNums, float64(nums[i]))
    }
    return judge(transNums)
}

func judge(nums []float64) bool {
    l := len(nums)

    if l == 0 {
        return false
    }

    if (l == 1) {
        return math.Abs(nums[0] - 24) < EPSILON
    }

    for i := 0; i < l; i ++ {
        for j := 0; j < l; j ++ {
            if i == j {
                continue
            }
            newNums := []float64{}
            for k := 0; k < l; k ++ {
                if k != i && k != j {
                    newNums = append(newNums, nums[k])
                }
            }

            if i < j && judge(append(newNums, nums[i] + nums[j])) {
                return true
            }
            if judge(append(newNums, nums[i] - nums[j])) {
                return true
            }
            if i < j && judge(append(newNums, nums[i] * nums[j])) {
                return true
            }
            if math.Abs(nums[j]) > EPSILON && judge(append(newNums, nums[i] / nums[j])) {
                return true
            }
        }
    } 
    return false
}