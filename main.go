package main

import "fmt"

func main() {

	// solange es nicht completed ist, versuch es zu completen, wenn keine züge mehr möglich sind fange von vorn an mit einem anderen seed.

	board := [8][8]bool{}

	currentX := 7
	currentY := 7
	board[currentX][currentY] = true

	for k := 0; k < 3; k++ {
		printCurrentBoard(board, currentX, currentY)
		fmt.Println()
		board, currentX, currentY = movePiece(board, currentX, currentY, -1, -2)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {

		}
	}

}

func movePiece(board [8][8]bool, currentX int, currentY int, x int, y int) ([8][8]bool, int, int) {
	fmt.Printf("Moving from (%d,%d) to (%d,%d)\n", currentX, currentY, currentX+x, currentY+y)
	board[currentX+x][currentY+y] = true
	return board, currentX + x, currentY + y
}

func printCurrentBoard(board [8][8]bool, currentX int, currentY int) {

	var ANSI_COLOR_RED string = "\x1b[31m"
	var ANSI_COLOR_BLUE string = "" //"\x1b[34m"
	var ANSI_COLOR_YELLOW string = "\x1b[33m"
	var ANSI_COLOR_RESET string = "\x1b[0m"

	// i and j are always swapped! i=y, j=x coordinates! might want to fix this
	// board[j][i] to access board at (x,y)
	for i := 0; i < 8; i++ {

		for j := 0; j < 8; j++ {

			currentPiece := ""

			if i == currentY && j == currentX {
				currentPiece = ANSI_COLOR_YELLOW + "@"
			} else if board[j][i] {
				currentPiece = "-"
			} else {
				currentPiece = "x"
			}

			if i%2 == 0 {
				j++
				if j%2 == 0 {
					fmt.Printf("%s%s%s ", ANSI_COLOR_RED, currentPiece, ANSI_COLOR_RESET)
				} else {
					fmt.Printf("%s%s%s ", ANSI_COLOR_BLUE, currentPiece, ANSI_COLOR_RESET)
				}
				j--
			} else {
				if j%2 == 0 {
					fmt.Printf("%s%s%s ", ANSI_COLOR_RED, currentPiece, ANSI_COLOR_RESET)
				} else {
					fmt.Printf("%s%s%s ", ANSI_COLOR_BLUE, currentPiece, ANSI_COLOR_RESET)
				}
			}
		}
		fmt.Println()
	}
}

func isBoardCompleted(board [8][8]bool) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !board[j][i] {
				return false
			}
		}
	}
	return true
}
