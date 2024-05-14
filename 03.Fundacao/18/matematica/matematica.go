package matematica

func Soma[T int | float64](a T, b T) T {
	return a + b
}
