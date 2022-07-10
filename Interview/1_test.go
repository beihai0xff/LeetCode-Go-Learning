package Interview

import "testing"

func Test_findMiddleMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal_test", args{nums: []int{21, 11, 45, 56, 9, 66, 77, 89, 78, 68, 100, 120, 111}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleMax(tt.args.nums); got != tt.want {
				t.Errorf("findMiddleMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
