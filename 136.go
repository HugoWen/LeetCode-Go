package main
import (
	"fmt"
)

func main() {
	var arr = []int{4,1,2,1,2}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
    ret := nums[0]
    if len(nums) > 1 {
        for i := 1; i < len(nums); i++ {
            ret = ret ^ nums[i]
        }
    }
    return ret
}