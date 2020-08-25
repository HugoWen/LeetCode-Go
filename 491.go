package main

import "strings"
import "fmt"

var hash map[string]bool
var ret [][]int

func findSubsequences(nums []int) [][]int {
    var hash map[string]bool
    var ret [][]int
    hash = make(map[string]bool)
    var dfs1 func(int, []int)
    dfs1 = func (start int, path []int) {
        if len(path) >= 2 {
            str := strings.Replace(strings.Trim(fmt.Sprint(path), "[]"), " ", ",", -1)
            _, ok := hash[str]
            if false == ok {
                tmp := make([]int, len(path))
                copy(tmp, path)
                ret = append(ret, tmp)
                hash[str] = true
            }
        }

        for i := start; i < len(nums); i ++ {
            if len(path) == 0 || nums[i] >= path[len(path) - 1] {
                path = append(path, nums[i])
                dfs1(i + 1, path)
                path = path[:(len(path) - 1)]
            }
        }
    }
    dfs1(0, []int{})
    return ret
}