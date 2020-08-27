package main

import "fmt"

func main() {
	// var arr = []int{4,6,7,7}
	// fmt.Println(findSubsequences(arr))
	// var str = "23"
	// fmt.Println(letterCombinations(str))

	var matrix = [][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,50},
	}
	// matrix = [][]int{
	// 	{1,3,5},
	// }
	fmt.Println(searchMatrix(matrix, 3))
}