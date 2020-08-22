func twoSum(nums []int, target int) []int {
    l := len(nums)
    ret := []int{}
    if l == 0 {
        return ret
    }

    hash := make(map[int]int, l)
    for i := 0; i < l; i ++ {
        minus := target - nums[i]
        pos, find := hash[nums[i]]
        if find {
            ret = append(ret, pos, i)
            return ret
        }
        hash[minus] = i
    }
    return ret
}