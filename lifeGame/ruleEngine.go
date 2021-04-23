package lifeGame

type RuleEngine struct {
}

func (this *RuleEngine)trastition(cellState CellState, numberOfNeighbours int) (CellState) {
	switch {
	case numberOfNeighbours == 2:
		return cellState
	case numberOfNeighbours == 3:
		return Live
	case numberOfNeighbours < 2:
		return Dead
	case numberOfNeighbours > 3:
		return Dead
	default:
		return cellState
	}
}
