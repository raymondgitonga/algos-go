package main

func treeSum(root *NodeInt) int {
	if root == nil {
		return 0
	}
	return root.Key + treeSum(root.Right) + treeSum(root.Left)
}
