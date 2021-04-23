package lifeGame

import (
	"github.com/ahmetb/go-linq"
	"fmt"
	"sync"
)

type Board struct {
	cells      [10][10]Cell
	turns      int
	dimension  int
	ruleEngine RuleEngine
	cacheCells [10][10]Cell
	channel  chan int
	wg sync.WaitGroup
}

func NewBoard(dimension int, turns int) *Board {
	cells := createCells(dimension)
	ruleEngine := RuleEngine{}
	channel := make(chan int,0)
	wg := sync.WaitGroup{}
	return &Board{
		cells: cells,
		turns: turns,
		dimension:dimension,
		ruleEngine:ruleEngine,
		channel:channel,
		wg: wg,
	}
}

func createCells(_ int)[10][10]Cell{
	cells := [10][10]Cell{}
	cells[0][0] = Cell{Dead}
	cells[0][1] = Cell{Live}
	cells[0][2] = Cell{Dead}
	cells[1][0] = Cell{Dead}
	cells[1][1] = Cell{Live}
	cells[1][2] = Cell{Dead}
	cells[2][0] = Cell{Dead}
	cells[2][1] = Cell{Live}
	cells[2][2] = Cell{Dead}
	return cells
}

func (this *Board) Play() {
	for i := 1; i <= this.turns; i++ {
		this.wg.Add(this.dimension*this.dimension)
		this.turn()
		this.wg.Wait()
	}
}

func (this *Board)Display() {
	for i := 0; i < this.dimension; i++ {
		fmt.Print("\n")
		for j := 0; j < this.dimension; j++ {
			fmt.Print(this.cells[i][j].displayString(), " ")
		}
	}
}

func (this *Board) turn() {
	this.cacheCells = this.cells
	for i := 0; i < this.dimension; i++ {
		for j := 0; j < this.dimension; j++ {
			go this.handleCell(i,j)
		}
	}
}

func (this *Board) handleCell(row int, column int) {
	numberOfNeighbours := this.getLiveNeighboursCount(row, column)
	this.cells[row][column].state = this.ruleEngine.trastition(this.cells[row][column].state, numberOfNeighbours)
	defer this.wg.Done()
}

func (this *Board) getLiveNeighboursCount(row int, column int) int {
	neighBourCells := []Cell{}
	if row <= 0 && column <= 0 {
		cell := this.cacheCells[row+1][column+1];
		neighBourCells = append(neighBourCells, cell)
		cell = this.cacheCells[row+1][column];
		neighBourCells = append(neighBourCells, cell)
		cell = this.cacheCells[row][column+1];
		neighBourCells = append(neighBourCells, cell)
	}
	if row > 0 && column > 0 {
		cell := this.cacheCells[row-1][column-1];
		neighBourCells = append(neighBourCells, cell)
	}
	if row > 0 {
		cell := this.cacheCells[row-1][column]
		neighBourCells = append(neighBourCells, cell)
		cell = this.cacheCells[row][column+1]
		neighBourCells = append(neighBourCells, cell)
	}
	if column > 0{
		cell := this.cacheCells[row][column-1]
		neighBourCells = append(neighBourCells, cell)
		cell = this.cacheCells[row+1][column]
		neighBourCells = append(neighBourCells, cell)
	}
	return linq.From(neighBourCells).Where(func(cell interface{}) bool {
		return cell.(Cell).state == Live
	}).Count()
}
