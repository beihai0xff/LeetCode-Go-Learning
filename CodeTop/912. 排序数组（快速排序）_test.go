package CodeTop

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
