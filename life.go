package main

import "fmt"
import "time"

var grid [20][20]rune

func readGrid() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println()
	}
}

func initGrid() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			grid[i][j] = '.'
		}
	}
}

func setPos(x int, y int) {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if i == x && j == y {
				grid[i][j] = 'X'
			}

		}
	}
}

func DecideFate(corX int, corY int) bool {
	var count int = -1
	if grid[corX-1][corY-1] == 'X' {
		count += 1
	}

	if grid[corX-1][corY] == 'X' {
		count += 1
	}

	if grid[corX][corY-1] == 'X' {
		count += 1
	}

	if grid[corX][corY-1] == 'X' {
		count += 1
	}

	if grid[corX+1][corY] == 'X' {
		count += 1
	}

	if grid[corX][corY+1] == 'X' {
		count += 1
	}

	if grid[corX+1][corY+1] == 'X' {
		count += 1
	}

	return count >= 3
}

func updateGrid() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if DecideFate(i, j) {
				grid[i][j] = '.'
			} else {
				grid[i][j] = 'X'
			}
		}
	}
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	initGrid()
	setPos(5, 5)
	setPos(5, 6)
	setPos(6, 6)
	setPos(4, 6)
	readGrid()
	if DecideFate(5, 6) {
		fmt.Println("Should Live")
	}
	time.Sleep(5)
	ClearTerminal()
	updateGrid()
	readGrid()
}
