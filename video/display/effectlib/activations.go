package effectlib

import "math"

// rms is the RMS of the audio signal, between 0 and 1.
// these groups of functions changes the value of rms
// the outputs are still between 0 and 1

func clamp(x float64) float64 {
	return math.Min(1, math.Max(0, x))
}

func Linear(x float64) float64 {
	return x
}

func Quadratic(x float64) float64 {
	return x * x
}

func Cubic(x float64) float64 {
	return x * x * x
}

func Smoothstep(x float64) float64 {
	return x * x * (3 - 2*x)
}

func Smootherstep(x float64) float64 {
	return x * x * x * (10 - 15*x + 6*x*x)
}

func Sine(x float64) float64 {
	return math.Sin(x * 0.5 * math.Pi)
}

func Truncatedlinear(x float64) float64 {
	return clamp(1.5*x - 0.25)
}

func MoreTruncatedlinear(x float64) float64 {
	return clamp(2*x - 0.5)
}
