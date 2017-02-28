// Author: Toufiqur Chowdhury
// Author URL: http://linkedin.com/in/toufiq
//

// Simulates a simple toy robot movement in a 2 dimensional space.
//
// What is ToyRobotSimulator?
//
// Toy robot simulator as the name suggest is an simplest robot command-line based movement
// simulation tool written in Go Language.
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var iRobot Robot
var iSurface Surface


/*
 * main (re-)starts the simulator and takes the length of the square-sized Surface
 * from user input in the console and generate an Adjacency Matrix
 * to be used as the surface to simulate robotic movements.
 */
func main() {
	iRobot = Robot{}
	iSurface = Surface{}
	fmt.Print("\nPlease enter the length of the adjacency matrix to start simulation: \n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		length, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err == nil && length > 1 {
			//initialize adjacency matrix
			iSurface = NewSurface(length)
			break
		}
		fmt.Println("Please enter the length of the adjacency matrix: >1 \n")
	}

	iSurface.Print(Node{})
	beginSimulation()
}

/*
 * beginSimulation starts accepting console input commands from user.
 * This following commands are accepted: (case-insensitive)
 *  PLACE X,Y,F  (X int, Y int, f string: east, west, north or south)
 *  MOVE
 *  LEFT
 *  RIGHT
 *  REPORT
 *  PRINT
 *  EXIT
 */
func beginSimulation() {
	fmt.Print("Enter command: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		processCommand(scanner.Text())
		fmt.Print("Enter command: ")
	}
}

const commandList = "Please enter one of the following commands: (case-insensitive)" +
	"\nPLACE X,Y,FrontDirection" +
	"\nMOVE" +
	"\nLEFT" +
	"\nRIGHT" +
	"\nREPORT" +
	"\nPRINT" +
	"\nRESET" +
	"\nEXIT"
/*
 * Process Robots command received from the console
 * Params:
 * @command string
 */
func processCommand(command string) {
	var err error
	parts := strings.Split(command, " ")
	switch(strings.ToLower(parts[0])) {
	default:
		//ignore unsupported command
		fmt.Println(commandList)
		break;
	case "place":
		//place robot on the surface/map
		//read x,y and f
		if len(parts) <= 1 {
			fmt.Println("Missing required arguments: X,Y,F")
			break
		}

		commandArgs := strings.Split(parts[1], ",")
		if len(commandArgs) < 3 {
			fmt.Println("Missing one or more required arguments: X,Y,F")
			break
		}

		x, _ := strconv.Atoi(strings.TrimSpace(commandArgs[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(commandArgs[1]))
		front := strings.TrimSpace(commandArgs[2])
		err = iRobot.Place(iSurface, x, y, front)
		break
	case "move":
		//move robot one step forward
		err = iRobot.Move(iSurface)
		break
	case "left":
		//change direction of the robot to left
		err = iRobot.Left(iSurface)
		break;
	case "right":
		//change direction of the robot to right
		err = iRobot.Right(iSurface)
		break;
	case "report":
		//print robots current position and direction
		iRobot.Report()
		iSurface.Print(iRobot.Node)
		break
	case "print":
		//print map (matrix) showing current position
		// of the robot (symbolized by character 1)
		iSurface.Print(iRobot.Node)
		break
	case "reset":
		main()
	case "exit":
		os.Exit(0)
		break
	}

	// Print any error/message
	if err != nil {
		fmt.Println(err)
	}
}