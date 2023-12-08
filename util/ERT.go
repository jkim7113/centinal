package util

import (
	"math"
	"strings"
)

func EstimateReadingTime(str string) int {
	wordLength := len(strings.Fields(str))
	ERT := math.Round(float64(wordLength) / 200.0)
	return int(math.Max(ERT, 1))
}
