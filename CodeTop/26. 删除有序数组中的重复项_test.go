package CodeTop

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 1, 1, 2, 4, 4, 6, 7}}, []int{1, 2, 4, 6, 7, 1, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeDuplicates2(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("removeDuplicates2() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
