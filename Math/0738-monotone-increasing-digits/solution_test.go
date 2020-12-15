/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/15 下午9:04
 * Author: beihai
 */

package _738_monotone_increasing_digits

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test0", args{322}, 299},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.N); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
