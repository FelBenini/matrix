package matrix

func Lerp[T Number](u, v Matrix[T], t T) Matrix[T] {
	res := Copy(u)
	v_cpy := Copy(v)
	v_cpy.Sub(res)
	v_cpy.Scl(t)
	res.Add(v_cpy)
	return (res)
}
