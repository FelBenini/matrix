package vector

func Lerp[T Number](u, v Vector[T], t T) Vector[T] {
	res := Vector[T]{
		Data: append([]T(nil), u.Data...),
	}
	v_sub := Vector[T]{
		Data: append([]T(nil), v.Data...),
	}
	v_sub.Sub(res)
	v_sub.Scl(t)
	res.Add(v_sub)
	return res
}
