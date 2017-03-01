package trs

// ExampleProcessCommand tests user input commands to manipulate robot movement
func ExampleProcessCommand() {
	iSurface = NewSurface(10)
	iRobot = Robot{}

	ProcessCommand("PLACE 7,8,SOUTH")
	ProcessCommand("MOVE")
	ProcessCommand("MOVE")
	ProcessCommand("REPORT")

	// Output:
	// 7,6,south
	//   -------------------- Y
	// [ 0 0 0 0 0 0 0 0 0 0 ] 9
	// [ 0 0 0 0 0 0 0 0 0 0 ] 8
	// [ 0 0 0 0 0 0 0 0 0 0 ] 7
	// [ 0 0 0 0 0 0 0 1 0 0 ] 6
	// [ 0 0 0 0 0 0 0 0 0 0 ] 5
	// [ 0 0 0 0 0 0 0 0 0 0 ] 4
	// [ 0 0 0 0 0 0 0 0 0 0 ] 3
	// [ 0 0 0 0 0 0 0 0 0 0 ] 2
	// [ 0 0 0 0 0 0 0 0 0 0 ] 1
	// [ 0 0 0 0 0 0 0 0 0 0 ] 0
	//   --------------------
	// X 0 1 2 3 4 5 6 7 8 9 X

}
