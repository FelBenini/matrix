package matrix

import (
	"fmt"
	"github.com/FelBenini/matrix/srcs/number"
)

type Number = number.Number

type Matrix[T Number] struct {
	Rows int
	Cols int
	Data [][]T
}

func From[T Number](data [][]T) Matrix[T] {
	rows := len(data)
	cols := len(data[0])

	copyData := make([][]T, rows)
	for i := range data {
		copyData[i] = append([]T(nil), data[i]...)
	}
	return Matrix[T]{
		Rows: rows,
		Cols: cols,
		Data: copyData,
	}
}

func Copy[T Number](m Matrix[T]) Matrix[T] {
	res := m

	copyData := make([][]T, m.Rows)
	for i := range m.Data {
		copyData[i] = append([]T(nil), m.Data[i]...)
	}
	res.Data = copyData
	return res
}

func (m Matrix[T]) Print() {
	fmt.Print("[")
	for i := range m.Rows {
		if i > 0 {
			fmt.Print(" ")
		}
		for j := range m.Cols {
			if m.Data[i][j] > 0 {
				fmt.Print(" ")
			}
			fmt.Print(m.Data[i][j])
			if j < m.Cols-1 {
				fmt.Print(" ")
			}
		}
		if i < m.Rows-1 {
			fmt.Println()
		}
	}
	fmt.Println("]")
}
