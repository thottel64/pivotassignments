package calculator_test

import (
	"github.com/thottel64/pivotassignments/calculator"
	"testing"
)

type testingvariables struct {
	a               int
	b               int
	desired         int
	operation       func(int, int) int
	divideoperation func(int, int) (int, error)
	err             error
	poweroperation  func(float64, float64) float64
}

var testingmap = map[string]testingvariables{
	"Add":          {a: 3, b: 2, desired: 5, operation: calculator.Add},
	"Subtract":     {a: 10, b: 7, desired: 3, operation: calculator.Subtract},
	"Multiply":     {a: 6, b: 4, desired: 24, operation: calculator.Multiply},
	"Divide":       {a: 12, b: 3, desired: 4, divideoperation: calculator.Divide},
	"DivideByZero": {a: 1, b: 0, desired: 0, divideoperation: calculator.Divide, err: calculator.DivideByZero{}},
	"Power":        {a: 2, b: 3, desired: 8, poweroperation: calculator.Pow},
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
			if test.poweroperation != nil {
				result := test.poweroperation(float64(test.a), float64(test.b))
				if int(result) != test.desired {
					t.Errorf("got %v, want %v", result, test.desired)
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

		})
	}
}
