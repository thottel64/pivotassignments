package calculator_test

import (
	"github.com/thottel64/pivotassignments/calculator"
	"testing"
)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
type testingvariables struct {
	a               int
	b               int
	desired         int
	operation       func(int, int) int
	divideoperation func(int, int) (int, error)
	err             error
}

var testingmap = map[string]testingvariables{
	"Add":          {a: 3, b: 2, desired: 5, operation: calculator.Add},
	"Subtract":     {a: 10, b: 7, desired: 3, operation: calculator.Subtract},
	"Multiply":     {a: 6, b: 4, desired: 24, operation: calculator.Multiply},
	"Divide":       {a: 12, b: 3, desired: 4, divideoperation: calculator.Divide},
	"DivideByZero": {a: 1, b: 0, desired: 0, divideoperation: calculator.Divide, err: calculator.DivideByZero{}},
}

func TestFunctions(t *testing.T) {

	for name, test := range testingmap {
		t.Run(name, func(t *testing.T) {
			if test.operation != nil {
				result := test.operation(test.a, test.b)
				if result != test.desired {
					t.Errorf("got %v, want %v", result, test.desired)
					return
				}
			}
			if test.divideoperation != nil {
				result, err := test.divideoperation(test.a, test.b)
				if err != test.err {
					t.Errorf("Did not receieve expected error")
					return
				}
				if test.err == nil && err != nil {
					t.Errorf("received unexcpected error %q", err)
					return
				}
				if test.desired != result {
					t.Errorf("got %v, want %v", result, test.desired)
					return
				}
			}

=======
func CalulatorTest(t *testing.T) {
=======
func TestCalculator(t *testing.T) {
>>>>>>> 8f7567e (fixed error with naming conventions)
	tests := map[string]struct {
		a, b, want         int
		operation          func(int, int) int
		operationWithError func(int, int) (int, error)
		err                error
	}{
		"Add":          {a: 1, b: 2, want: 3, operation: calculator.Add},
		"Subtract":     {a: 1, b: 2, want: -1, operation: calculator.Subtract},
		"Multiply":     {a: 2, b: 3, want: 6, operation: calculator.Multiply},
		"Divide":       {a: 6, b: 3, want: 2, operationWithError: calculator.Divide},
		"DivideByZero": {a: 6, b: 0, want: 0, operationWithError: calculator.Divide, err: calculator.DivideByZero{}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.operation != nil {
				got := test.operation(test.a, test.b)
				if got != test.want {
					t.Errorf("got %d, want %d", got, test.want)
				}
				return
			}

			got, err := test.operationWithError(test.a, test.b)

			if test.err != nil {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if err.Error() != test.err.Error() {
					t.Errorf("got %q, want %q", err, test.err)
				}
			}

			if test.err == nil && err != nil {
				t.Errorf("got %q, want nil", err)
			}

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
>>>>>>> d270329 (added calculator funcs and init commit for testing)
		})
=======
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
=======
func TestFunctions(t *testing.T) {
	testmap := map[string]struct {
		a               int
		b               int
		desired         int
		operation       func(int, int) int
		divideoperation func(int, int) (int, error)
		err             error
	}{
		"Add":          {a: 3, b: 2, desired: 5, operation: calculator.Add},
		"Subtract":     {a: 10, b: 7, desired: 3, operation: calculator.Subtract},
		"Multiply":     {a: 6, b: 4, desired: 24, operation: calculator.Multiply},
		"Divide":       {a: 12, b: 3, desired: 4, divideoperation: calculator.Divide},
		"DivideByZero": {a: 1, b: 0, desired: 0, divideoperation: calculator.Divide, err: calculator.DivideByZero{}},
>>>>>>> d5428be (updated test functions to be table-driven)
	}
	for name, test := range testmap {
		t.Run(name, func(t *testing.T) {
			if test.operation != nil {
				result := test.operation(test.a, test.b)
				if result != test.desired {
					t.Errorf("got %v, want %v", result, test.desired)
					return
				}
			}
			if test.divideoperation != nil {
				result, err := test.divideoperation(test.a, test.b)
				if err != test.err {
					t.Errorf("Did not receieve expected error")
					return
				}
				if test.err == nil && err != nil {
					t.Errorf("received unexcpected error %q", err)
					return
				}
				if test.desired != result {
					t.Errorf("got %v, want %v", result, test.desired)
					return
				}
			}

<<<<<<< HEAD
}

func TestDivideByZero_Error(t *testing.T) {
	got, err := calculator.Divide(1, 0)
	if got != 0 || err == nil {
		t.Errorf("calculator is not handling divide by zero error correctly, no error received")
	}
	if err != error(calculator.DivideByZero{}) {
		t.Errorf("different erorr than expected given. Wanted %v, got %v", error(calculator.DivideByZero{}), err)
>>>>>>> e1165fb (reworked calculator tests as instructed)
=======
		})
>>>>>>> d5428be (updated test functions to be table-driven)
	}
}
