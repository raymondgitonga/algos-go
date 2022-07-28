package main

func treeIncludes(root *Node, target string) bool {
	if root == nil {
		return false
	}
	if root.Key == target {
		return true
	}

	if root.Right != nil {
		return treeIncludes(root.Right, target)
	}

	if root.Left != nil {
		return treeIncludes(root.Left, target)
	}

	return false
}
