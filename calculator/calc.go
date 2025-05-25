package calculator

type Calculator struct{}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
