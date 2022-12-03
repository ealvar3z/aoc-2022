package numbers

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[T number](input T) T {
	if input < 0 {
		return -input
	}
	return input
}

func Max[T number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
