/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/23 下午10:03
 * Author: beihai
 */

package _069_sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{4}, 2},
		{"test2", args{8}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrt1(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{4}, 2},
		{"test2", args{8}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt1(tt.args.x); got != tt.want {
				t.Errorf("mySqrt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_mySqrt(b *testing.B) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{4}, 2},
		{"test2", args{8}, 2},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mySqrt(tt.args.x)
		}
	}
}

func Benchmark_mySqrt1(b *testing.B) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{4}, 2},
		{"test2", args{8}, 2},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mySqrt1(tt.args.x)
		}
	}
}
