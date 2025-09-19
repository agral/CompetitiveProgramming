package lc3484

import (
	"strconv"
	"strings"
)

// Runtime complexity: O(1) for each query.
// Auxiliary space: O(|SetCell|).
// Subjective level: medium.
type Spreadsheet struct {
	Cells map[string]int
}

func Constructor(rows int) Spreadsheet {
	s := new(Spreadsheet)
	s.Cells = make(map[string]int)
	return *s
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.Cells[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.SetCell(cell, 0)
}

func (this *Spreadsheet) GetValue(formula string) int {
	parts := strings.Split(formula[1:], "+")
	return this.evaluate(parts[0]) + this.evaluate(parts[1])
}

func (this *Spreadsheet) evaluate(text string) int {
	if converted, err := strconv.Atoi(text); err == nil {
		return converted
	}
	value, isOk := this.Cells[text]
	if !isOk {
		return 0
	}
	return value
}
