package matrix

func (m Matrix[T]) Add(m2 Matrix[T]) {
	for i := range m.Rows {
		for j := range m.Cols {
			m.Data[i][j] += m2.Data[i][j]
		}
	}
}
