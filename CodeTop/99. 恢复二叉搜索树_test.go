package CodeTop

import "testing"

func Test_recoverTree(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
		Right: nil,
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{root}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
		})
	}
}
