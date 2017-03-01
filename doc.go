// Author: Toufiqur Chowdhury
// Author URL: http://linkedin.com/in/toufiq
//

// Simulates a simple toy robot movement in a 2-dimensional space.
//
// Introduction
//
// Toy robot simulator as the name suggest is an simplest robot command-line based movement
// simulation tool written in Go Language. It uses a 2-dimensional space using a square matrix
// to represent the surface.
//
// How to run application
//
// There are several ways you can run a Go Application.
// Such as compiling an executable file or running in the terminal.
// Easiest and quickest is to use the terminal.
// In either case, you will need to have Go installed.
// You can download and install it from here: https://golang.org/doc/install
// To run the application in the terminal,
// go to the ToyRobotSimulator directory in a terminal
// and run the following command:
//   go run *.go
//
// Supported Commands
//
// This following commands are accepted: (case-insensitive)
//  PLACE X,Y,F  // places robot on the surface coordinates :X int, Y int, f string (east, west, north or south)
//  MOVE 		 // moves the robot one step forward in the direction it is currently facing
//  LEFT		 // changes robot's direction to immediate left
//  RIGHT		 // changes robot's direction to immediate right
//  REPORT		 // outputs robot's coordinate X,Y and F direction. Also executes PRINT command
//  PRINT        // outputs the 2-dimensional surface matrix with X-Y axis information
//  RESET		 // resets all variables and restarts the simulation
//  EXIT		 // exits the application
//
// How to use
//
// Upon running the application, the user will be prompted to input a number to setup the square
// surface for simulation purposes. User can then input any of the above mentioned commands,
// but all actions will be prevented until a valid "PLACE" command is entered.
// Here's a series of example commands after surface has been set:
//   place 3,4,east
//   move
//   move
//   report
//
// Output:
//   5,4,east
//   -------------------- Y
//   [ 0 0 0 0 0 0 0 0 0 0 ] 9
//   [ 0 0 0 0 0 0 0 0 0 0 ] 8
//   [ 0 0 0 0 0 0 0 0 0 0 ] 7
//   [ 0 0 0 0 0 0 0 0 0 0 ] 6
//   [ 0 0 0 0 0 0 0 0 0 0 ] 5
//   [ 0 0 0 0 0 1 0 0 0 0 ] 4
//   [ 0 0 0 0 0 0 0 0 0 0 ] 3
//   [ 0 0 0 0 0 0 0 0 0 0 ] 2
//   [ 0 0 0 0 0 0 0 0 0 0 ] 1
//   [ 0 0 0 0 0 0 0 0 0 0 ] 0
//   --------------------
//   X 0 1 2 3 4 5 6 7 8 9 X
package main
