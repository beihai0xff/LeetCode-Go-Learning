/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/9 下午9:36
 * Author: beihai
 */

package _190_reverse_bits

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"test1", args{43261596}, 964176192},
		{"test2", args{4294967293}, 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBits1(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"test1", args{43261596}, 964176192},
		{"test2", args{4294967293}, 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits1(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseBits(b *testing.B) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"test1", args{43261596}, 964176192},
		{"test2", args{4294967293}, 3221225471},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reverseBits(tt.args.num)
		}
	}
}

func Benchmark_reverseBits1(b *testing.B) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"test1", args{43261596}, 964176192},
		{"test2", args{4294967293}, 3221225471},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reverseBits1(tt.args.num)
		}
	}
}
