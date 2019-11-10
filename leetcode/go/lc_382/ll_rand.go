package lc_382

import (
	"math/rand"
	"time"
)

// Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type Solution struct {
	vals []int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	ret := Solution{make([]int, 0, 64)}
	for ; head != nil; head = head.Next {
		ret.vals = append(ret.vals, head.Val)
	}
	rand.Seed(time.Now().Unix())
	return ret
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
