package angles

import "math"

const (
	Tau float64 = 2 * math.Pi
	Turn
	Degrees = Tau / 360
	Grads   = Tau / 400
)

// TODO: How to name the inverse transform?
// If r is in radians, what is the operation and the constant
// to coax a degrees out of it?
