package vector

import (
	"math"
)

func (v Vector[T]) Norm_1() T {
	var res T
	size := len(v.Data)
	res = 0
	for i := range size {
		res += v.Data[i]
	}
	return res
}

func (v Vector[T]) Norm() T {
	var res T
	size := len(v.Data)
	res = 0
	for i := range size {
		res += T(math.Pow(float64(v.Data[i]), 2))
	}
	res = T(math.Sqrt(float64(res)))
	return res
}

func (v Vector[T]) Norm_inf() T {
	var res T

	res = 0
	for _, val := range v.Data {
		abs := val
		if val < 0 {
			abs = -val
		}
		if abs > res {
			res = abs
		}
	}
	return res
}
