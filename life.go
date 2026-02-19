package main

import "fmt"

var grid [20][20]string

func readGrid() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("%s ", grid[i][j])
		}
		fmt.Println()
	}
}

func initGrid() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			grid[i][j] = "."
		}
	}
}

func setPos(x int, y int) {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if i == x && j == y {
				grid[i][j] = "X"
			}

		}
	}
}

func main() {
	initGrid()
	setPos(1, 1)
	setPos(1, 2)
	setPos(2, 2)
	readGrid()
}
