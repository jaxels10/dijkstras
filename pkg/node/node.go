package node

type Node struct {
	id int
	x  float64
	y  float64
}

func newNode(id int, x float64, y float64) *Node {
	return &Node{id: id, x: x, y: y}
}
