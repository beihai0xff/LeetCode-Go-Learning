package CodeTop

import "testing"

func Test_numOfx(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{12, 1}, 5},
		{"1", args{1234567, 1}, 1058985},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfx(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("numOfx() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_numOfx2(t *testing.T) {
// 	type args struct {
// 		n int
// 		x int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"1", args{12, 1}, 5},
// 		{"1", args{12, 2}, 2},
// 		{"1", args{1234567, 1}, 1058985},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := numOfx2(tt.args.n, tt.args.x); got != tt.want {
// 				t.Errorf("numOfx() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_countDigitOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1234567}, 1058985},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
