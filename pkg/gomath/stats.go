package gomath

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math"
	"sort"
)

func Mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	return goutils.SumBy(values, func(v float64) float64 {
		return v
	}) / float64(len(values))
}

func StandardDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	mean := Mean(numbers)
	sum := 0.0
	for _, n := range numbers {
		sum += (n - mean) * (n - mean)
	}
	return math.Sqrt(sum / float64(len(numbers)))
}

func Median(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}
