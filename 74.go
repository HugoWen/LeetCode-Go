package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    ret := false
    for i := len(matrix) - 1; i >= 0; i -- {
        if matrix[i][0] > target {
            continue;
        }

        if (matrix[i][0]) == target {
            return true
        }

		fmt.Println(i)
        return binarySearch(matrix[i], target)
    }
    return ret
}

func binarySearch(matrix []int, target int) bool {
    ret := false
    for {
		fmt.Println(matrix)
        if len(matrix) == 0 {
            return false
        }
        mid := len(matrix) / 2
        if matrix[mid] == target {
            return true
        } else if matrix[mid] > target {
            matrix = matrix[0: mid]
            return binarySearch(matrix, target)
        } else {
            matrix = matrix[mid + 1:]
            return binarySearch(matrix, target)
        }
    }
    return ret
}