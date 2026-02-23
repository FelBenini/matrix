package main

import (
	"fmt"
	"github.com/FelBenini/matrix/srcs/matrix"
	"github.com/FelBenini/matrix/srcs/vector"
)

func main() {
	fmt.Println("\033[1m----ex00----\033[0m")
	fmt.Println("\nVectors:")
	u := vector.From([]float32{2, 3})
	v := vector.From([]float32{5, 7})
	u.Add(v)
	u.Print()
	u = vector.From([]float32{2, 3})
	v = vector.From([]float32{5, 7})
	u.Sub(v)
	u.Print()
	u = vector.From([]float32{2, 3})
	u.Scl(2)
	u.Print()

	fmt.Println("\nMatrixes:")
	m := matrix.From([][]float32{{1, 2}, {3, 4}})
	n := matrix.From([][]float32{{7, 4}, {-2, 2}})
	m.Add(n)
	m.Print()
	m = matrix.From([][]float32{{1, 2}, {3, 4}})
	n = matrix.From([][]float32{{7, 4}, {-2, 2}})
	m.Sub(n)
	m.Print()
	m = matrix.From([][]float32{{1, 2}, {3, 4}})
	m.Scl(2)
	m.Print()
}
