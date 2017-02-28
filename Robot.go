package main

import (
	"errors"
	"fmt"
	"strings"
)

type Robot struct {
	Node
	Front string //front direction: east, west, north, south
}
/*
 * Place robot on the surface using x,y coordinates and front-direction.
 * If placement fails returns error.
 *
 * Params:
 * @x int : x-axis location
 * @y int : y-axis location
 * @f string: Robot's front direction.
 *            Accepted values: east, west, north, south
 *
 * Returns error
 */
func (robot *Robot) Place(surface Surface, x, y int, font string) (err error) {

	//validate front direction
	if surface.GetDirectionIndex(font) == -1 {
		return errors.New("Invalid front direction supplied!")
	}

	newNode := Node{x, y, false}
	//validate node positioning
	if err = surface.PlaceNode(&newNode); err != nil {
		return err
	}

	// success
	robot.Node = newNode
	robot.Front = strings.ToLower(font)
	return nil
}
/*
 * Move robot one step forward in the direction it is currently facing
 *
 * Params:
 * @surface Surface
 *
 * Returns error
 */
func (robot *Robot) Move(surface Surface) (err error) {
	var nextNode Node
	nextNode, err = surface.GetNextNodeByDirection(robot.Node, robot.Front)
	if err != nil {
		return
	}

	robot.Node = nextNode
	return
}

func (robot *Robot) Face(surface Surface, head string) (error) {
	if !robot.Placed {
		return errors.New("Robot is not placed yet!")
	}
	lastIndex := len(ValidDirections) - 1
	index := surface.GetDirectionIndex(robot.Front)
	if index == -1 {
		return errors.New("Robot has an invalid current direction.")
	}
	nextIndex := index - 1
	if head != "left" {// !left, assume right
		nextIndex = index + 1
	}
	if nextIndex < 0 {
		nextIndex = lastIndex
	} else if nextIndex > lastIndex {
		nextIndex = 0
	}
	robot.Front = ValidDirections[nextIndex]
	return nil
}

/*
 * Change robot's direction to immediate left of current direction
 */
func (robot *Robot) Left(surface Surface) (error) {
	return robot.Face(surface, "left")
}

/*
 * Change robot's direction to immediate right of current direction
 */
func (robot *Robot) Right(surface Surface) (error) {
	return robot.Face(surface, "right")
}

/*
 * Report announces current status of Robot.
 * Robot has been place only if robot.F has a value.
 * In that case will return a message stating robots position and direction (f) and nil error
 * Otherwise, will return an error and empty message.
 *
 * Returns: Node, string, error
 */
func (robot *Robot) Report() {

	if robot.Front == "" || !robot.Placed {
		fmt.Println("Robot has not been placed yet!")
		return
	}

	fmt.Println(fmt.Sprint(robot.X) + "," + fmt.Sprint(robot.Y) + "," + robot.Front)
}