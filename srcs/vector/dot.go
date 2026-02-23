package vector

func (v Vector[T]) Dot(v2 Vector[T]) T {
	var res T
	res = 0
	for i := range len(v.Data) {
		res += v.Data[i] * v2.Data[i]
	}
	return (res)
}
