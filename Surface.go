package main

import (
	"strings"
	"errors"
	"fmt"
)

type Surface struct {
	Nodes           []Node
	AdjacencyMatrix map[int] []int
}

/*
 * NewSurface creates a surface with a square area, represented by an adjacency matrix
 *
 *  Params:
 *  @length int : width and length of the surface
 *
 *  Returns Surface
 */
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

/*
 * ValidDirections lists supported front directions for any robot
 */
var ValidDirections = []string{
	"north", //y-axis +
	"east", //x-axis +
	"south", //y-axis -
	"west", //x-axis -
}

/*
 * Print adjacency matrix to visually represent the surface and a node position
 */
func (surface *Surface) Print(node Node) {
	line := ""
	for l := 0; l < len(iSurface.AdjacencyMatrix) * 2; l++ {
		line += "-"
	}
	fmt.Println("  " + line + " Y")

	for x := 0; x < len(iSurface.AdjacencyMatrix); x++ {
		y := len(iSurface.AdjacencyMatrix) - x -1;
		fmt.Print("[ ")
		for x2 := 0; x2 < len(iSurface.AdjacencyMatrix); x2++ {
			val := fmt.Sprint(iSurface.AdjacencyMatrix[x2][y])
			if node.Placed && x2 == node.X && y == node.Y {
				val = "1"
			}
			fmt.Print(val, " ")
		}
		fmt.Print("] ", y,"\n")
	}

	fmt.Println("  " + line)
	fmt.Print("X ")

	for i := 0; i < len(iSurface.AdjacencyMatrix[0]); i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("X")
}

/*
 * Checks if supplied string is an acceptable direction listed in ValidDirections
 *
 * Params:
 * @d string
 *
 * Returns bool
 */
func (surface *Surface) GetDirectionIndex(front string) (int) {
	front = strings.ToLower(front)
	for index, dir := range ValidDirections {
		if front == strings.ToLower(dir) {
			return index
		}
	}
	return -1
}

/*
 * Validate a nodes coordinates
 */
func (surface *Surface) PlaceNode(node *Node) (error) {
	//validate x-axis value
	if node.X < 0 || node.X >= len(surface.AdjacencyMatrix) {
		return errors.New("Invalid/unsupported x-axis value supplied!")
	}

	//validate y-axis value
	if node.Y < 0 || node.Y >= len(surface.AdjacencyMatrix[0]) {
		return errors.New("Invalid/unsupported x-axis value supplied!")
	}
	node.Placed = true
	return nil
}

/*
 * Move to the node forward
 * Params:
 * @node Node
 * @front string: front-facing direction
 *
 * Returns Node, error
 */
func (surface *Surface) GetNextNodeByDirection(node Node, front string) (nextNode Node, err error) {
	if !node.Placed {
		err = errors.New("Robot has not been placed yet!")
		return
	}
	nextNode = Node{node.X, node.Y, false}
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

	if err = surface.PlaceNode(&nextNode); err != nil {
		err = errors.New("You have reached the end of line! " +
			"\nChange direction by using 'LEFT' or 'RIGHT' command.")
		return
	}
	//success
	return
}
