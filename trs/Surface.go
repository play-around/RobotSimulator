package trs

import (
	"strings"
	"errors"
	"fmt"
)

// Surface struct is used to describe a 2 dimensional space and other relevant information
type Surface struct {
	Nodes           []Node // ToDo: can be used to place and track multiple robots
	AdjacencyMatrix map[int] []int
}

// NewSurface constructs an instance of the Surface struct
// and generates a square matrix using supplied number.
//
//  Params:
//  @length int
//
//  Returns Surface
func NewSurface(length int) (surface Surface){
	surface.AdjacencyMatrix = map[int] []int{}
	for i := 0; i < length; i++ {
		row := []int{}
		for j := 0; j < length; j++ {
			row = append(row, 0)
		}
		surface.AdjacencyMatrix[i] = row
	}
	return
}

// ValidDirections lists directions supported for robot's movement
var ValidDirections = []string{
	"north", //y-axis +1
	"east", //x-axis +1
	"south", //y-axis -1
	"west", //x-axis -1
}

// Print outputs the 2d representation of the surface with X,Y axis.
// If a Node is supplied as argument, it will be displayed on the matrix
// and it's position will be presented using '1' where as all other
// empty nodes will be presented using '0'
//
//  Params:
//  @node Node
func (surface *Surface) Print(node Node) {
	line := ""
	for l := 0; l < len(surface.AdjacencyMatrix) * 2; l++ {
		line += "-"
	}
	fmt.Println("  " + line + " Y")

	for x := 0; x < len(surface.AdjacencyMatrix); x++ {
		y := len(surface.AdjacencyMatrix) - x -1;
		fmt.Print("[ ")
		for x2 := 0; x2 < len(surface.AdjacencyMatrix); x2++ {
			val := fmt.Sprint(surface.AdjacencyMatrix[x2][y])
			if node.Placed && x2 == node.X && y == node.Y {
				val = "1"
			}
			fmt.Print(val, " ")
		}
		fmt.Print("] ", y,"\n")
	}

	fmt.Println("  " + line)
	fmt.Print("X ")

	for i := 0; i < len(surface.AdjacencyMatrix[0]); i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("X")
}

// IsValidDirection checks if supplied string is an acceptable direction listed in ValidDirections
//
//  Params:
//  @d string
//
//  Returns int
func (surface *Surface) GetDirectionIndex(d string) (int) {
	d = strings.ToLower(d)
	for index, dir := range ValidDirections {
		if d == strings.ToLower(dir) {
			return index
		}
	}
	return -1
}

// ValidateNode validates whether a node's coordinates are within the surface
//
//  Params:
//  @node Node
//
//  Returns error
func (surface *Surface) ValidateNode(node Node) (error) {
	//validate x-axis value
	if node.X < 0 || node.X >= len(surface.AdjacencyMatrix) {
		return errors.New("Invalid/unsupported x-axis value supplied!")
	}

	//validate y-axis value
	if node.Y < 0 || node.Y >= len(surface.AdjacencyMatrix[0]) {
		return errors.New("Invalid/unsupported x-axis value supplied!")
	}

	return nil
}

// GetNextNodeByDirection moves a robot one one step forward
//
//  Params:
//  @node Node
//  @front string: front-facing direction
//
//  Returns Node, error
func (surface *Surface) GetNextNodeByDirection(node Node, front string) (nextNode Node, err error) {
	if !node.Placed {
		err = errors.New("Robot has not been placed yet!")
		return
	}
	nextNode = NewNode(node.X, node.Y, false)
	switch (strings.ToLower(front)) {
	case "east":
		nextNode.X++
		break
	case "west":
		nextNode.X--
		break
	case "north":
		nextNode.Y++
		break
	case "south":
		nextNode.Y--
		break
	default:
		err = errors.New("Invalid direction supplied!")
		return
	}

	if err = surface.ValidateNode(nextNode); err != nil {
		err = errors.New("You have reached the end of line!" +
			"\nChange direction by using 'LEFT' or 'RIGHT' command.")
		return
	}
	// success
	nextNode.Placed = true
	return
}
