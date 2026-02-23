package number

func Lerp[T Number](u, v, t T) T {
	res := (v - u) * t
	res += u
	return (res)
}
