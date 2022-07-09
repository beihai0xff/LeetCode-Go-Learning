package CodeTop

import "testing"

func Test_rotateLeft(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{"test", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}}},
		{"test", [][]int{{1}}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateLeft(tt.matrix)
		})
	}
}
