package CodeTop

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	var helper func(nums []int, left, right int) *TreeNode
	helper = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(nums, left, mid-1)
		root.Right = helper(nums, mid+1, right)
		return root
	}

	return helper(nums, 0, len(nums)-1)
}
