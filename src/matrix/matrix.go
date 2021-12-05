package matrix

import (
	"fmt"

	"github.com/ThomVett/aoc2021/src/utils"
)

const height = 5
const width = 5

type IntMatrix2D struct {
	Data     [][]int
	Selected [][]bool
	Name     string
}

func (m IntMatrix2D) Contains(number int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if number == m.Data[i][j] {
				m.Selected[i][j] = true
				fmt.Println(m.Selected)
			}
		}
	}
}

func (m *IntMatrix2D) WinningRow() bool {

	for i := 0; i < height; i++ {
		if utils.SumBool(m.Selected[i][:]) == width {
			return true
		}
	}
	return false
}

func (m *IntMatrix2D) WinningCol() bool {
	for i := 0; i < height; i++ {
		col_selections := []bool{}
		for j := 0; j < width; j++ {
			col_selections = append(col_selections, m.Selected[i][j])
		}
		if utils.SumBool(m.Selected[i][:]) == width {
			return true
		}
	}
	return false
}
