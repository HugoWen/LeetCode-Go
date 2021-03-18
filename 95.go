package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return _generateTrees(1, n)
}

func _generateTrees(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	trees := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := _generateTrees(start, i-1)
		right := _generateTrees(i+1, end)
		for _, l := range left {
			for _, r := range right {
				cur := &TreeNode{i, nil, nil}
				cur.Left = l
				cur.Right = r
				trees = append(trees, cur)
			}
		}
	}
	return trees
}
