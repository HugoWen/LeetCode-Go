package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var buildTree func(int, int) *TreeNode
	buildTree = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		node := &TreeNode{
			Val: 0,
			Left: nil,
			Right: nil,
		}
		mid := (left + right) / 2
		node.Val = nums[mid]
		node.Left = buildTree(left, mid - 1)
		node.Right = buildTree(mid + 1, right)
		return node
	}

	return buildTree(0, len(nums) - 1)
}