package matrix

func (m Matrix[T]) Scl(s T) {
	for i := range m.Rows {
		for j := range m.Cols {
			m.Data[i][j] *= s
		}
	}
}
