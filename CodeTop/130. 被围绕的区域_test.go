package CodeTop

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"1", args{[][]byte{[]byte("XXXX"), []byte("XOOX"), []byte("XXOX"), []byte("XOXX")}}, [][]byte{[]byte("XXXX"), []byte("XXXX"), []byte("XXXX"), []byte("XOXX")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
