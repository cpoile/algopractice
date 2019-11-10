package lc_21

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = &ListNode{Val: l2.Val}
			cur = cur.Next
			l2 = l2.Next
		} else if l2 == nil {
			cur.Next = &ListNode{Val: l1.Val}
			cur = cur.Next
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			cur = cur.Next
			l2 = l2.Next
		}
	}
	return head.Next
}
