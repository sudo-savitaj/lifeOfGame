package lifeGame

type CellState int

const (
	Dead CellState = 0
	Live CellState = 1
)

type Cell struct {
	state CellState;
}

func NewCell(state CellState) *Cell {
	return &Cell{
		state: 0,
	}
}

func (this *Cell)displayString() string{
	if this.state == Live {
		return "X"
	} else {
		return "-"
	}
}
