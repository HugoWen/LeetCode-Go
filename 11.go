package main

func maxArea(height []int) int {
    area := 0
    i := 0 
    j := len(height) - 1
    for ;i < j;  {
        area = max(area, min(height[i], height[j]) * (j - i))
        if height[i] <= height[j] {
            i ++
        } else {
            j --
        }
    }
    return area
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}