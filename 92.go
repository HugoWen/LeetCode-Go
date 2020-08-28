package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head.Next == nil || m >= n {
        return head
    }

    newHead := &ListNode{
        Val: 0,
        Next: head,
    }
    pre := newHead

    for i := 0; i < m - 1; i ++ {
        pre = pre.Next
    }

    cur := pre.Next
    for i := 0; i < n - m; i ++ {
        tmp := pre.Next
        pre.Next = cur.Next
        cur.Next = cur.Next.Next
        pre.Next.Next = tmp
    }

    return newHead.Next
}