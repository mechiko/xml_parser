package address

import (
	"strconv"
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
	colChar := string(a.col + 65)
	out := colChar + strconv.Itoa(a.row)
	return out
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
	colChar := string(a.col + col + 65)
	out := colChar + strconv.Itoa(a.row)
	return out
}

func (a *Address) Range(row int, col int) string {
	colChar := string(a.col + col + 65)
	out := colChar + strconv.Itoa(a.row+row)
	return out
}

func (a *Address) Row() int {
	return a.row
}

func (a *Address) Col() int {
	return a.row
}
