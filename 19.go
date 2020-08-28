package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || n <= 0 {
        return head
    }

    newHead := &ListNode{
        Val: 0,
        Next: head,
    }
    first := newHead
    second := newHead

    for i := 0; i <= n; i ++ {
        first = first.Next
    }

    for first != nil {
        first = first.Next
        second = second.Next
    }

    second.Next = second.Next.Next
    return newHead.Next
}