package CodeTop

import "testing"

func Test_isPalindrome125(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"Aa"}, true},
		{"2", args{"ab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome125(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome125() = %v, want %v", got, tt.want)
			}
		})
	}
}
