package trs

import (
	"errors"
	"fmt"
	"strings"
)

// Robot is use to store information a Robot
type Robot struct {
	Node
	Front string //front direction: east, west, north, south
}
// NewRobot constructs an instance of Robot
func NewRobot(x, y int, front string) (robot Robot){
	robot.X = x
	robot.Y = y
	robot.Front = front
	return
}

// Place robot on the surface using x,y coordinates and front-direction.
// If placement fails returns error.
//
//  Params:
//  @x int : x-axis location
//  @y int : y-axis location
//  @f string: Robot's front direction. Accepted values: east, west, north, south
//
//  Returns error
func (robot *Robot) Place(surface Surface, x, y int, f string) (err error) {

	//validate front direction
	if surface.GetDirectionIndex(f) == -1 {
		return errors.New("Invalid front direction supplied!")
	}

	//validate node positioning
	if err = surface.ValidateNode(robot.Node); err != nil {
		return err
	}

	robot.Node = NewNode(x, y, true)
	robot.Front = strings.ToLower(f)
	return nil
}

// Move changes robot's position one step forward in the direction it is currently facing
//
//  Params:
//  @surface Surface
//
//  Returns error
//
func (robot *Robot) Move(surface Surface) (err error) {
	var nextNode Node
	nextNode, err = surface.GetNextNodeByDirection(robot.Node, robot.Front)
	if err != nil {
		return
	}

	robot.Node = nextNode
	return
}

/*
 * Face changes robot's direction to immediate left/right of current direction.
 *
 *  Params:
 *  @surface Surface
 *  @side string : left/right
 *
 *  Returns error
 */
func (robot *Robot) ChangeDirection(surface Surface, side string) (error) {
	if !robot.Placed {
		return errors.New("Robot is not placed yet!")
	}
	lastIndex := len(ValidDirections) - 1
	index := surface.GetDirectionIndex(robot.Front)
	if index == -1 {
		return errors.New("Robot has an invalid current direction.")
	}
	nextIndex := index - 1
	if side != "left" {// !left, assume right
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
 * Left changes robot's direction to immediate left of current direction
 *
 *  Params:
 *  @surface Surface
 *
 *  Returns error
 */
func (robot *Robot) Left(surface Surface) (error) {
	return robot.ChangeDirection(surface, "left")
}

/*
 * Right changes robot's direction to immediate right of current direction
 *
 *  Params:
 *  @surface Surface
 *
 *  Returns error
 */
func (robot *Robot) Right(surface Surface) (error) {
	return robot.ChangeDirection(surface, "right")
}

// Report prints current coordinates and directions of the robot.
// Robot has been place only if robot.F has a value.
// In that case will return a message stating robots position and direction (f) and nil error
// Otherwise, will return an error and empty message.
//
//  Returns: string
//
func (robot *Robot) Report() (string) {
	if robot.Front == "" || !robot.Placed {
		return "Robot has not been placed yet!"
	}

	return fmt.Sprint(robot.X) + "," + fmt.Sprint(robot.Y) + "," + robot.Front
}