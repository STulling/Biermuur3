package math

import (
	"math"
)

func ProcessBlock(block []int16) (float64, float64) {
	return calcRMS(block), 0
}

func calcRMS(block []int16) float64 {
	var sum float64 = 0
	for _, sample := range block {
		sum += float64(sample) / 32768.0
	}
	return math.Sqrt(sum / float64(len(block)))
}
