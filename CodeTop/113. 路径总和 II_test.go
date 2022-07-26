package CodeTop

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   7,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   2,
				},
				Val: 11,
			},
			Right: nil,
			Val:   4,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   13,
			},
			Right: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   5,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   1,
				},
				Val: 4,
			},
			Val: 8,
		},
		Val: 5,
	}
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test_normal", args{root: root, targetSum: 22}, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
