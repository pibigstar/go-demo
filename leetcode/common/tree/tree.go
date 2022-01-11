package tree

type TreeNode struct {
	Val   int
	Depth int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func GetTree() *TreeNode {
	node8 := NewTreeNode(8, nil, nil)
	node7 := NewTreeNode(7, nil, nil)
	node6 := NewTreeNode(6, node7, nil)
	node5 := NewTreeNode(5, nil, nil)
	node4 := NewTreeNode(4, node8, nil)
	node3 := NewTreeNode(3, node6, nil)
	node2 := NewTreeNode(2, node4, node5)
	node1 := NewTreeNode(1, node2, node3)
	return node1
}

func GetHaveDuplicateTree() *TreeNode {
	node444 := NewTreeNode(4, nil, nil)
	node44 := NewTreeNode(4, nil, nil)
	node22 := NewTreeNode(2, node44, nil)
	node4 := NewTreeNode(4, nil, nil)
	node3 := NewTreeNode(3, node22, node444)
	node2 := NewTreeNode(2, node4, nil)
	node1 := NewTreeNode(1, node2, node3)
	return node1
}
