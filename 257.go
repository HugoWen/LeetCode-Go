package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }

    if root.Left == nil && root.Right == nil {
        return []string{strconv.Itoa(root.Val)}
    }
    
    var ret []string    
    var recur func(*TreeNode, string)
    recur = func(root *TreeNode, s string) {
        if s != "" {
            s += "->" + strconv.Itoa(root.Val)
        } else {
            s = strconv.Itoa(root.Val)
        }
        if root.Left == nil && root.Right == nil {
            ret = append(ret, s)
            return
        }
        if root.Left != nil {
            recur(root.Left, s)
        }
        if root.Right != nil {
            recur(root.Right, s)
        }
    }
    recur(root, "")
    return ret
}