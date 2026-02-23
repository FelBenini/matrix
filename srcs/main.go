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
	fmt.Println("\n\033[1m----ex01----\033[0m")
	v1 := vector.From([]float32{1, 0, 0})
	v2 := vector.From([]float32{0, 1, 0})
	v3 := vector.From([]float32{0, 0, 1})
	v_res := vector.LinearCombination([]vector.Vector[float32]{v1, v2, v3}, []float32{10, -2, 0.5})
	v_res.Print()
	v1 = vector.From([]float32{1, 2, 3})
	v2 = vector.From([]float32{0, 10, -100})
	v_res = vector.LinearCombination([]vector.Vector[float32]{v1, v2}, []float32{10, -2})
	v_res.Print()
}
