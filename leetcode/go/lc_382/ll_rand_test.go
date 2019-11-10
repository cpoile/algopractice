package lc_382

import (
	"fmt"
	"testing"
)

func NewLL(ints []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, v := range ints {
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
	}

	return head.Next
}

func TestSolution_GetRandom(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want int
	}{
		{
			name: "first",
			list: []int{1, 2, 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor(NewLL(tt.list))
			for i := 0; i < 10; i++ {
				fmt.Printf("%v, ", obj.GetRandom())
			}
			//if got := obj.GetRandom(); got != tt.want {
			//	t.Errorf("Solution.GetRandom() = %v, want %v", got, tt.want)
			//}
		})
	}
}
