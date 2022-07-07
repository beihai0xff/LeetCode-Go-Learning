package CodeTop

import "testing"

func Test_isValid(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"test", "()[]{}", true},
		{"test", "(]", false},
		{"test", "{[]}", true},
		{"test", "(])", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
