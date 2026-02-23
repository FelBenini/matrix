package vector

import (
	"math"
)

func (v Vector[T]) Norm_1() T {
	var res T
	res = 0
	for _, val := range v.Data {
		res += val
	}
	return res
}

func (v Vector[T]) Norm() T {
	var res T
	res = 0
	for _, val := range v.Data {
		res += T(math.Pow(float64(val), 2))
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
