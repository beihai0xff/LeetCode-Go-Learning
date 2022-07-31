package CodeTop

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{0}, []int{1}},
		{"2", args{1}, []int{1, 1}},
		{"3", args{2}, []int{1, 2, 1}},
		{"4", args{3}, []int{1, 3, 3, 1}},
		{"5", args{4}, []int{1, 4, 6, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
