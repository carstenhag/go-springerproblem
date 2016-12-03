package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

const s = 8 // Board size. 5x5, 6x6 are instant, 7x7 takes 30-120 seconds, 8x8 ???

func main() {

	// Launch as many threads as cores available
	for i := 0; i < runtime.NumCPU(); i++ {
		go algo()
	}

	// Impedes main() to end. Goroutines run in the background.
	fmt.Println("Calculating...")
	fmt.Scanln()

}

func algo() {
	start := time.Now()
	board := [s][s]bool{}

	// Selbst hier brauchen wir schon einen neuen Seed. -- bestimmt falsch
	seed := int64(time.Now().UnixNano())
	rnd := rand.New(rand.NewSource(seed))

	initX := rnd.Intn(s)
	initY := rnd.Intn(s)

	currentX := initX
	currentY := initY
	board[currentX][currentY] = true

	for {

		seed := int64(time.Now().UnixNano()) + rand.Int63()
		rndSeed := rand.NewSource(seed)
		rnd := rand.New(rndSeed)

		currentX = rand.Intn(s)
		currentY = rand.Intn(s)

		/*
			// Playback, geht noch nicht
			rnd = rand.New(rand.NewSource(-8228114560393910474))
			currentX = 2
			currentY = 0
		*/

		// Upping the number of times (200--> 300) reduced 8x8 time from 20-140 minutes to 90 seconds.
		for k := 0; k < 300; k++ {
			r := 0
			if !isBoardCompleted(board) {

				// Debug:
				//printCurrentBoard(board, currentX, currentY)
				//time.Sleep(200 * time.Millisecond)

				r = rnd.Intn(7 + 1)
				switch r {
				case 0:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, 2, 1)
					}
				case 1:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, 2, -1)
					}
				case 2:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, -2, 1)
					}
				case 3:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, -2, -1)
					}
				case 4:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, 1, 2)
					}
				case 5:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, 1, -2)
					}
				case 6:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, -1, 2)
					}
				case 7:
					{
						board, currentX, currentY = movePiece(board, currentX, currentY, -1, -2)
					}
				}

			} else {
				fmt.Printf("\x1b[32mSize: %dx%d, Elapsed Time: %s\n", s, s, time.Since(start).String())
				fmt.Printf("Seed: %d, Initial X: %d, Initial Y: %d\n", seed, initX, initY)
				//fmt.Println("\x1b[32m" + "DONE" + "\x1b[0m")
				printCurrentBoard(board, currentX, currentY)
				os.Exit(0)
			}

		}
		// Reinitialize board to 0
		board = [s][s]bool{}
	}
}

func movePiece(board [s][s]bool, currentX int, currentY int, x int, y int) ([s][s]bool, int, int) {
	if currentX+x >= 0 && currentX+x <= s-1 && currentY+y >= 0 && currentY+y <= s-1 {
		if !board[currentX+x][currentY+y] {
			//fmt.Printf("Moving from (%d,%d) to (%d,%d)\n", currentX, currentY, currentX+x, currentY+y)
			board[currentX+x][currentY+y] = true
			return board, currentX + x, currentY + y
		}
	}
	return board, currentX, currentY
}

func printCurrentBoard(board [s][s]bool, currentX int, currentY int) {

	var ANSI_COLOR_RED string = "\x1b[31m"
	var ANSI_COLOR_BLUE string = "" //"\x1b[34m" //Left blank because it's pretty ugly
	var ANSI_COLOR_YELLOW string = "\x1b[33m"
	var ANSI_COLOR_RESET string = "\x1b[0m"

	// i and j are always swapped! i=y, j=x coordinates! might want to fix this
	// board[j][i] to access board at (x,y)
	for i := 0; i < s; i++ {

		for j := 0; j < s; j++ {

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
	fmt.Println()
}

func isBoardCompleted(board [s][s]bool) bool {
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if !board[j][i] {
				return false
			}
		}
	}
	return true
}
