package audiomath

import (
	"math"
)

func ProcessBlock(block []float32) [2]float64 {
	result := [2]float64{0, 0}
	result[0] = calcRMS(block)
	return result
}

func calcRMS(block []float32) float64 {
	var sum float64
	for _, sample := range block {
		sum += math.Pow(float64(sample), 2)
	}
	return math.Sqrt(sum / float64(len(block)))
}
