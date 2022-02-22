package list

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int, next *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
	}
}

// 1 -> 2 -> 3 -> 4 -> 5
func GetHead() *Node {
	node5 := NewNode(5, nil)
	node4 := NewNode(4, node5)
	node3 := NewNode(3, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	return node1
}

// 0 -> 6 -> 7
func GetHead2() *Node {
	node7 := NewNode(7, nil)
	node6 := NewNode(6, node7)
	node0 := NewNode(0, node6)
	return node0
}

// 1 -> 2 -> 2 -> 1 -> nil
func GetPalindromes() *Node {
	node4 := NewNode(1, nil)
	node3 := NewNode(2, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	return node1
}

// 1 -> 2 -> 3 -> 1
func GetCycle() *Node {
	var node1 *Node
	node3 := NewNode(3, nil)
	node2 := NewNode(2, node3)
	node1 = NewNode(1, node2)

	node3.Next = node1
	return node1
}
