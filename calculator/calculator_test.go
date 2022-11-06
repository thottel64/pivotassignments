package calculator_test

import (
	"github.com/thottel64/pivotassignments/calculator"
	"testing"
)

func TestCalculator(t *testing.T) {
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
		})
	}
}
