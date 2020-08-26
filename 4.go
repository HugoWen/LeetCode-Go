package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l := len(nums1) + len(nums2)
    mid, mid1 := 0, 0
    if l % 2 == 1 {
        mid = l / 2
        return float64(findK(nums1, nums2, mid + 1))
    } else {
        mid = l / 2 - 1
        mid1 = l / 2
        // todo: optimize
        return float64((findK(nums1, nums2, mid + 1) + findK(nums1, nums2, mid1 + 1))) / 2.0
    }
}

func findK(nums1 []int, nums2[]int, k int) int {
    index1, index2 := 0, 0
    for {
        if index1 == len(nums1) {
            return nums2[index2 + k - 1]
        }

        if index2 == len(nums2) {
            return nums1[index1 + k - 1]
        }

        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }

        half := k / 2
        newIndex1 := min(index1 + half - 1, len(nums1) - 1)
        newIndex2 := min(index2 + half - 1, len(nums2) - 1)
        if nums1[newIndex1] >= nums2[newIndex2] {
            k -= (newIndex2 - index2 + 1)
            index2 = newIndex2 + 1
        } else {
            k -= (newIndex1 - index1 + 1)         
            index1 = newIndex1 + 1
        }
    }
    
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}