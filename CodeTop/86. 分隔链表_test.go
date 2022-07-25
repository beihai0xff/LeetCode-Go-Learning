package CodeTop

import (
	"reflect"
	"testing"
)

func NewSingleList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummyHead := &ListNode{}
	curr := dummyHead
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return dummyHead.Next
}

func Test_partition(t *testing.T) {

	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{head: NewSingleList([]int{1, 4, 3, 2, 5, 2}), x: 3}, []int{1, 2, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValues(partition(tt.args.head, tt.args.x)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
