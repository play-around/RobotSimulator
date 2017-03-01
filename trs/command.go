package trs

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "errors"
)

// iRobot is an instance of Robot
var iRobot Robot

// iSurface is an instance of Surface
var iSurface Surface

// SetupSimulation takes the length of the square-sized Surface
// from user input in the console and generate an Adjacency Matrix
// to be used as the surface to simulate robotic movements.
func SetupSimulation() {
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
    BeginSimulation()
}

// BeginSimulation starts accepting console input commands from user.
func BeginSimulation() {
    fmt.Print("Enter command: \n")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ProcessCommand(scanner.Text())
        fmt.Print("Enter command: \n")
    }
}

// ProcessCommand executes a user command received from the console
//
//  Params:
//  @command string
func ProcessCommand(command string) {
    var err error
    parts := strings.Split(command, " ")
    switch(strings.ToLower(parts[0])) {
    default:
        //ignore unsupported command
        err = errors.New("Supported commands: (case-insensitive)"+
            "\nPLACE X,Y,F"+
            "\nMOVE"+
            "\nLEFT"+
            "\nRIGHT"+
            "\nREPORT"+
            "\nPRINT"+
            "\nRESET"+
            "\nEXIT")
        break;
    case "place":
        //place robot on the surface/map
        //read x,y and f
        if len(parts) <= 1 {
            err = errors.New("Missing required arguments: X,Y,F")
            break
        }

        commandArgs := strings.Split(parts[1], ",")
        if len(commandArgs) < 3 {
            err = errors.New("Missing one or more required arguments: X,Y,F")
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
        fmt.Println(iRobot.Report())
        iSurface.Print(iRobot.Node)
        break
    case "print":
        //print map (matrix) showing current position
        // of the robot (symbolized by character 1)
        iSurface.Print(iRobot.Node)
        break
    case "reset":
        SetupSimulation()
    case "exit":
        os.Exit(0)
        break
    }

    if err != nil {
        fmt.Println(err)
    }
}
