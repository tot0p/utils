package tree

// Node represents a node in a binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode creates a new Node with the given value
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// Insert adds a new value to the binary tree
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = NewNode(value)
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(value)
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search looks for a value in the binary tree
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}
	if value < n.Value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}
