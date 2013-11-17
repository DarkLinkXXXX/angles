package angles

import (
	"math"
	"testing"
)

func equals(x, y float64) bool {
	return math.Abs(x-y) < 1e-10
}

func Test(t *testing.T) {
	var tests = []struct {
		expected, actual float64
	}{
		{360 * Degrees, 1 * Turn},
		{200 * Grads, math.Pi},
		{math.Sin(Tau / 4), 1},
	}

	for k, v := range tests {
		if !equals(v.expected, v.actual) {
			t.Log("Expected", v.expected)
			t.Log("Actual", v.actual)
			t.Error("Test", k, "failed.")
		}
	}
}
