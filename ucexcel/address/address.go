package address

import (
	"github.com/xuri/excelize/v2"
)

type Address struct {
	row int
	col int
}

func New(row int, col int) *Address {
	return &Address{
		row: row,
		col: col,
	}
}

func (a *Address) Address() string {
	// excelize expects 1-based col/row
	name, _ := excelize.CoordinatesToCellName(a.col+1, a.row)
	return name
}

func (a *Address) NextCol() string {
	a.col += 1
	out := a.Address()
	return out
}

func (a *Address) NextRow() string {
	a.row += 1
	a.col = 0
	out := a.Address()
	return out
}

func (a *Address) MoveTo(row int, col int) string {
	a.row = row
	a.col = col
	out := a.Address()
	return out
}
func (a *Address) AddCol(col int) string {
	a.col += col
	out := a.Address()
	return out
}
func (a *Address) AddRow(row int) string {
	a.row += row
	out := a.Address()
	return out
}

func (a *Address) ShiftCol(col int) string {
	name, _ := excelize.CoordinatesToCellName(a.col+1+col, a.row)
	return name
}

func (a *Address) Range(row int, col int) string {
	name, _ := excelize.CoordinatesToCellName(a.col+1+col, a.row+row)
	return name
}

func (a *Address) Row() int {
	return a.row
}

func (a *Address) Col() int {
	return a.col
}
