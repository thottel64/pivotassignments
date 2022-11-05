package calculator_test

import (
	"github.com/thottel64/pivotassignments/calculator"
	"testing"
)

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
	}
}
