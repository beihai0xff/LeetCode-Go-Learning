package CodeTop

import "testing"

func Test_reverse7(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{123}, 321},
		{"1", args{1534236469}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse7(tt.args.x); got != tt.want {
				t.Errorf("reverse7() = %v, want %v", got, tt.want)
			}
		})
	}
}
