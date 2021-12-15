package Offer

import (
	"reflect"
	"testing"
)

func Test_exchange(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"case1", []int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
