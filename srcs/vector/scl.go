package vector

func (v Vector[T]) Scl(s T) {
	count := len(v.Data)

	for i := range count {
		v.Data[i] *= s
	}
}
