package lc_21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		l1, l2   *ListNode
		expected *ListNode
	}{
		{&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}},
	}

	for _, c := range cases {
		actual := mergeTwoLists(c.l1, c.l2)
		assert.Equal(t, c.expected, actual, "given %v, %v", c.l1, c.l2)
	}
}
