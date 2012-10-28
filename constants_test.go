package angles

import (
	"math"
	"testing"
)

func equals(x, y float64) bool {
	return math.Abs(x-y) < 1e-20
}

func TestTurns(t *testing.T) {
	if !equals(360*Degrees, 1*Turn) {
		t.Fail()
	}
}
