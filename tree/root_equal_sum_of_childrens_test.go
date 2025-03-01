package leetcode

import "testing"

func TestRootEqualSumsOfChildren(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{name: "Valid tree", root: &TreeNode{Val: 10, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}, want: true},
		{name: "Invalid tree", root: &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkTree(tt.root)
			if got != tt.want {
				t.Errorf("checkTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
