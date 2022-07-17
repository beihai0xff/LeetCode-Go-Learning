package CodeTop

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{
			l1: nil,
			l2: nil,
		}, nil},
		{"test", args{
			l1: l1,
			l2: l2,
		}, []int{1, 1, 3, 4, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValues(mergeTwoLists(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
