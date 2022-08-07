package CodeTop

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name     string
		args     args
		wantOnes int
	}{
		{"1", args{0x00000000000000000000000000001011}, 3},
		{"2", args{0x00000000000000000000000010000000}, 1},
		{"3", args{0x00000000000000000000000000000001}, 1},
		{"4", args{0x00000000000000000000000000000010}, 1},
		{"5", args{0x00000000000000000000000000000011}, 2},
		{"6", args{0x00000000000000000000000000000100}, 1},
		{"7", args{0x00000000000000000000000000000101}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOnes := hammingWeight(tt.args.num); gotOnes != tt.wantOnes {
				t.Errorf("hammingWeight() = %v, want %v", gotOnes, tt.wantOnes)
			}
		})
	}
}
