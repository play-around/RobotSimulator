package trs

import "fmt"

// Tests both the NewSurface() function.
// First a Surface object with 10 x 10 matrix is instantiated.
// Then the Surface is printed along with X and Y axis information using the Surface.Print() function
// A new Node with coordinates is supplied as an argument.
// Since the node is set as placed by setting true on the Surface,
// it will be represented inside the matrix and represented with '1'.
func ExampleNewSurface() {
	length := 10
	newSurface := NewSurface(length)
	node := NewNode(4,5, true)
	newSurface.Print(node)

	// Output:
	// -------------------- Y
	// [ 0 0 0 0 0 0 0 0 0 0 ] 9
	// [ 0 0 0 0 0 0 0 0 0 0 ] 8
	// [ 0 0 0 0 0 0 0 0 0 0 ] 7
	// [ 0 0 0 0 0 0 0 0 0 0 ] 6
	// [ 0 0 0 0 1 0 0 0 0 0 ] 5
	// [ 0 0 0 0 0 0 0 0 0 0 ] 4
	// [ 0 0 0 0 0 0 0 0 0 0 ] 3
	// [ 0 0 0 0 0 0 0 0 0 0 ] 2
	// [ 0 0 0 0 0 0 0 0 0 0 ] 1
	// [ 0 0 0 0 0 0 0 0 0 0 ] 0
	//   --------------------
	// X 0 1 2 3 4 5 6 7 8 9 X
}

// Tests the Surface.Print() function.
// First a Surface object with 10 x 10 matrix is instantiated.
// Then the Surface is printed along with X and Y axis information using the Surface.Print() function
// A new Node is supplied as an argument.
// Since the Node has not been placed on the Surface yet, it will not be displayed.
func ExampleSurface_Print() {
	length := 10
	newSurface := NewSurface(length)
	newSurface.Print(Node{})

	// Output:
	// -------------------- Y
	// [ 0 0 0 0 0 0 0 0 0 0 ] 9
	// [ 0 0 0 0 0 0 0 0 0 0 ] 8
	// [ 0 0 0 0 0 0 0 0 0 0 ] 7
	// [ 0 0 0 0 0 0 0 0 0 0 ] 6
	// [ 0 0 0 0 0 0 0 0 0 0 ] 5
	// [ 0 0 0 0 0 0 0 0 0 0 ] 4
	// [ 0 0 0 0 0 0 0 0 0 0 ] 3
	// [ 0 0 0 0 0 0 0 0 0 0 ] 2
	// [ 0 0 0 0 0 0 0 0 0 0 ] 1
	// [ 0 0 0 0 0 0 0 0 0 0 ] 0
	//   --------------------
	// X 0 1 2 3 4 5 6 7 8 9 X
}

// Tests the Surface.GetNextNodeByDirection() function.
// First a Surface object with 10 x 10 matrix is instantiated.
// Then the Surface is printed along with X and Y axis information using the Surface.Print() function
// A new Node is supplied as an argument.
// Since the Node has not been placed on the Surface yet, it will not be displayed.
func ExampleSurface_GetNextNodeByDirection() {
	length := 10
	newSurface := NewSurface(length)
	node := NewNode(4,5, true)
	nextNode, _ := newSurface.GetNextNodeByDirection(node, "east")
	newSurface.Print(nextNode)

	// Output:
	// -------------------- Y
	// [ 0 0 0 0 0 0 0 0 0 0 ] 9
	// [ 0 0 0 0 0 0 0 0 0 0 ] 8
	// [ 0 0 0 0 0 0 0 0 0 0 ] 7
	// [ 0 0 0 0 0 0 0 0 0 0 ] 6
	// [ 0 0 0 0 0 1 0 0 0 0 ] 5
	// [ 0 0 0 0 0 0 0 0 0 0 ] 4
	// [ 0 0 0 0 0 0 0 0 0 0 ] 3
	// [ 0 0 0 0 0 0 0 0 0 0 ] 2
	// [ 0 0 0 0 0 0 0 0 0 0 ] 1
	// [ 0 0 0 0 0 0 0 0 0 0 ] 0
	//   --------------------
	// X 0 1 2 3 4 5 6 7 8 9 X
}

// Tests of node is within inside the surface
func ExampleSurface_ValidateNode() {
	surface := NewSurface(10)
	node := NewNode(15, 5, true)
	if err := surface.ValidateNode(node); err != nil {
		fmt.Println(err)
	}

	// Output: Invalid/unsupported x-axis value supplied!
}

