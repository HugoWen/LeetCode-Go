package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    mid := len(lists) / 2
    left := mergeKLists(lists[:mid])
    right := mergeKLists(lists[mid:])
    return mergeTwoList(left, right)
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    if l1.Val <= l2.Val {
        l1.Next = mergeTwoList(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoList(l1, l2.Next)
        return l2
    }
}