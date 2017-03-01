package trs

import "fmt"

// Test reporting of robots status
func ExampleRobot_Report() {
	robot := NewRobot(6, 7, "north")
	fmt.Println(robot.Report())

	// Output: Robot has not been placed yet!
}

// Tests placing a robot on the surface
func ExampleRobot_Place() {
	surface := NewSurface(10)
	robot := Robot{}
	robot.Place(surface, 5, 6, "north")
	fmt.Println(robot.Report())

	// Output: 5,6,north
}

// ExampleMove tests placing a
// robot on the surface
func ExampleRobot_Move() {
	surface := NewSurface(10)
	robot := Robot{}
	robot.Place(surface, 5, 6, "north")
	robot.Move(surface)
	fmt.Println(robot.Report())

	// Output: 5,7,north
}

// Tests placing a robot reaching the end of the surface
func ExampleRobot_Move_hitEnd() {
	surface := NewSurface(10)
	robot := Robot{}
	robot.Place(surface, 5, 8, "north")

	//move robot forward
	err := robot.Move(surface)
	if err != nil {
		fmt.Println(err)
	}

	//move robot forward again
	err = robot.Move(surface)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// You have reached the end of line!
	// Change direction by using 'LEFT' or 'RIGHT' command.
}
