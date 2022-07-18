package CodeTop

import "testing"

func Test_maxSubArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test", args{[]int{1, 2, 3, 4, 5}}, 0, 4},
		{"test2", args{[]int{1, -2, 3, 4, 5}}, 2, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := maxSubArray2(tt.args.nums)
			if got != tt.want {
				t.Errorf("maxSubArray2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("maxSubArray2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
