package vector

func LinearCombination[T Number](v []Vector[T], coefs []T) Vector[T] {
	vector_len := len(v)
	size := len(v[0].Data)
	result := Vector[T]{Data: make([]T, size)}

	for i := range vector_len {
		for j := range size {
			result.Data[j] += v[i].Data[j] * coefs[i]
		}
	}
	return (result)
}
