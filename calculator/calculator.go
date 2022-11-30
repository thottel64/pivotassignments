package calculator

func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}
func Multiply(a int, b int) int {
	return a * b
}
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, DivideByZero{}
	}
	return a / b, nil
}

type DivideByZero struct{}

func (DivideByZero) Error() string {
	return "cannot divide by zero"
}
