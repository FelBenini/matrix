package vector

func (v Vector[T]) Sub(v2 Vector[T]) {
	count := len(v.Data)

	for i := range count {
		v.Data[i] -= v2.Data[i]
	}
}
