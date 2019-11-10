package lc_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	retHead := new(ListNode)
	retCur := retHead

	var carry int
	for l1 != nil || l2 != nil {
		var l1val, l2val int
		if l1 == nil {
			l1val = 0
		} else {
			l1val = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			l2val = 0
		} else {
			l2val = l2.Val
			l2 = l2.Next
		}

		sum := l1val + l2val + carry
		retCur.Next = &ListNode{Val: sum % 10}
		retCur = retCur.Next
		carry = sum / 10
	}
	if carry > 0 {
		retCur.Next = &ListNode{Val: carry}

		return retHead.Next
	}
}
