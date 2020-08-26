package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return addNum(l1, l2, false)
}

func addNum(l1 *ListNode, l2 *ListNode, carry bool) *ListNode {
    if l1 == nil && l2 == nil && carry == false {
        return nil
    }

    sum := 0
    if carry == true {
        sum ++
    } 

    if l1 != nil {
        sum += l1.Val
        l1 = l1.Next
    }

    if l2 != nil {
        sum += l2.Val
        l2 = l2.Next
    }

    node := ListNode{
        Val: sum % 10,
        Next: addNum(l1, l2, sum >= 10),
    }

    return &node
}