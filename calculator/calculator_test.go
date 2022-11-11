package calculator_test

import (
	"github.com/thottel64/pivotassignments/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	got := calculator.Add(3, 2)
	want := 5
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestSubtract(t *testing.T) {
	got := calculator.Subtract(10, 7)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestMultiply(t *testing.T) {
	got := calculator.Multiply(4, 6)
	want := 24
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestDivide(t *testing.T) {
	got, err := calculator.Divide(12, 3)
	if err != nil {
		t.Errorf("Divide function is returning an error")
	}
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestDivideByZero_Error(t *testing.T) {
	got, err := calculator.Divide(1, 0)
	if got != 0 || err == nil {
		t.Errorf("calculator is not handling divide by zero error correctly, no error received")
	}
	if err != error(calculator.DivideByZero{}) {
		t.Errorf("different erorr than expected given. Wanted %v, got %v", error(calculator.DivideByZero{}), err)
	}
}
