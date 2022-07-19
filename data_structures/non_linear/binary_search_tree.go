package non_linear

// Tree with no more than two children, child smaller than the parent
// is always on the left and larger on the right

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(k int) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{
				Key: k,
			}
		} else {
			n.Right.insert(k)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &Node{
				Key: k,
			}
		} else {
			n.Left.insert(k)
		}
	}
}

func (n *Node) search(k int) bool {
	if n == nil {
		return false
	}
	if k > n.Key {
		return n.Right.search(k)
	} else if k < n.Key {
		return n.Left.search(k)
	}

	return true
}
