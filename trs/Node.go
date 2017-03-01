package trs
//
// Node
//
// Node is used to store X,Y coordinates on a surface
type Node struct {
	X      int
	Y      int
	Placed bool
}

// Construct a new Node
//
//  Params:
//  @x int
//  @y int
//  @placed bool
//
// Returns Node
func NewNode(x, y int, placed bool) (node Node) {
	node.X = x
	node.Y = y
	node.Placed = placed
	return
}