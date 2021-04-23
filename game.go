package main

import (
	"main/lifeGame"
	"fmt"
)

func main()  {
	board := lifeGame.NewBoard(3,1)
	fmt.Println("-----------------------")
	fmt.Println("Game of Life")
	fmt.Println("-----------------------")
	board.Display()
	board.Play()
	fmt.Print("\n\n")
	board.Display()
}

